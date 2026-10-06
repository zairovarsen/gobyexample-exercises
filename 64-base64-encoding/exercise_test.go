package base64encoding

import "testing"

// Pairs of plain text and its standard / URL-safe base64 form. The two
// forms only differ where the alphabet does: "+" and "/" versus "-" and "_".
var samples = []struct {
	plain, std, url string
}{
	{"abc123!?$*&()'-=@~", "YWJjMTIzIT8kKiYoKSctPUB+", "YWJjMTIzIT8kKiYoKSctPUB-"},
	{">>>???", "Pj4+Pz8/", "Pj4-Pz8_"},
	{"hello", "aGVsbG8=", "aGVsbG8="},
	{"Go by Example", "R28gYnkgRXhhbXBsZQ==", "R28gYnkgRXhhbXBsZQ=="},
	{"a", "YQ==", "YQ=="},
	{"ab", "YWI=", "YWI="},
}

func TestEncode(t *testing.T) {
	for _, s := range samples {
		if got := Encode(s.plain); got != s.std {
			t.Errorf("Encode(%q) = %q, want %q", s.plain, got, s.std)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, s := range samples {
		got, err := Decode(s.std)
		if err != nil {
			t.Errorf("Decode(%q) returned error %v, want nil", s.std, err)
		}
		if got != s.plain {
			t.Errorf("Decode(%q) = %q, want %q", s.std, got, s.plain)
		}
	}
}

func TestDecodeErrors(t *testing.T) {
	tests := []struct {
		why     string
		encoded string
	}{
		{"it contains characters outside the alphabet", "not base64!"},
		{"it uses the URL-safe alphabet", "Pj4-Pz8_"},
		{"the padding is missing", "YQ"},
	}
	for _, tt := range tests {
		got, err := Decode(tt.encoded)
		if err == nil {
			t.Errorf("Decode(%q) returned a nil error, want an error because %s", tt.encoded, tt.why)
		}
		if got != "" {
			t.Errorf("Decode(%q) = %q together with an error, want an empty string", tt.encoded, got)
		}
	}
}

func TestEncodeURL(t *testing.T) {
	for _, s := range samples {
		if got := EncodeURL(s.plain); got != s.url {
			t.Errorf("EncodeURL(%q) = %q, want %q", s.plain, got, s.url)
		}
	}
}

func TestDecodeURL(t *testing.T) {
	for _, s := range samples {
		got, err := DecodeURL(s.url)
		if err != nil {
			t.Errorf("DecodeURL(%q) returned error %v, want nil", s.url, err)
		}
		if got != s.plain {
			t.Errorf("DecodeURL(%q) = %q, want %q", s.url, got, s.plain)
		}
	}

	for _, encoded := range []string{"Pj4+Pz8/", "not base64!", "YQ"} {
		got, err := DecodeURL(encoded)
		if err == nil {
			t.Errorf("DecodeURL(%q) returned a nil error, want an error (not valid URL-safe base64)", encoded)
		}
		if got != "" {
			t.Errorf("DecodeURL(%q) = %q together with an error, want an empty string", encoded, got)
		}
	}
}

func TestToURLSafe(t *testing.T) {
	for _, s := range samples {
		got, err := ToURLSafe(s.std)
		if err != nil {
			t.Errorf("ToURLSafe(%q) returned error %v, want nil", s.std, err)
		}
		if got != s.url {
			t.Errorf("ToURLSafe(%q) = %q, want %q", s.std, got, s.url)
		}
	}

	for _, std := range []string{"Pj4-Pz8_", "not base64!", "YQ"} {
		got, err := ToURLSafe(std)
		if err == nil {
			t.Errorf("ToURLSafe(%q) returned a nil error, want an error (not valid standard base64)", std)
		}
		if got != "" {
			t.Errorf("ToURLSafe(%q) = %q together with an error, want an empty string", std, got)
		}
	}
}

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		user, password string
		want           string
	}{
		{"aladdin", "open sesame", "Basic YWxhZGRpbjpvcGVuIHNlc2FtZQ=="},
		{"user", "p:ss", "Basic dXNlcjpwOnNz"},
		{"", "", "Basic Og=="},
	}
	for _, tt := range tests {
		if got := BasicAuth(tt.user, tt.password); got != tt.want {
			t.Errorf("BasicAuth(%q, %q) = %q, want %q", tt.user, tt.password, got, tt.want)
		}
	}
}

func TestParseBasicAuth(t *testing.T) {
	tests := []struct {
		header             string
		wantUser, wantPass string
	}{
		{"Basic YWxhZGRpbjpvcGVuIHNlc2FtZQ==", "aladdin", "open sesame"},
		{"Basic dXNlcjpwOnNz", "user", "p:ss"},
		{"Basic Og==", "", ""},
	}
	for _, tt := range tests {
		user, pass, err := ParseBasicAuth(tt.header)
		if err != nil {
			t.Errorf("ParseBasicAuth(%q) returned error %v, want nil", tt.header, err)
		}
		if user != tt.wantUser || pass != tt.wantPass {
			t.Errorf("ParseBasicAuth(%q) = %q, %q; want %q, %q", tt.header, user, pass, tt.wantUser, tt.wantPass)
		}
	}
}

func TestParseBasicAuthErrors(t *testing.T) {
	tests := []struct {
		why    string
		header string
	}{
		{`it does not start with "Basic "`, "Bearer YWxhZGRpbjpvcGVuIHNlc2FtZQ=="},
		{`it does not start with "Basic "`, "YWxhZGRpbjpvcGVuIHNlc2FtZQ=="},
		{"the credentials are not valid base64", "Basic not base64!"},
		{`the decoded credentials ("nocolon") contain no ":"`, "Basic bm9jb2xvbg=="},
		{"it is empty", ""},
	}
	for _, tt := range tests {
		user, pass, err := ParseBasicAuth(tt.header)
		if err == nil {
			t.Errorf("ParseBasicAuth(%q) returned a nil error, want an error because %s", tt.header, tt.why)
		}
		if user != "" || pass != "" {
			t.Errorf("ParseBasicAuth(%q) = %q, %q together with an error, want empty strings", tt.header, user, pass)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	for _, plain := range []string{"", "x", "héllo wörld ✓", "line one\nline two\n", "\x00\x01\xfe\xff"} {
		if got, err := Decode(Encode(plain)); err != nil || got != plain {
			t.Errorf("Decode(Encode(%q)) = %q, %v; want the original string and a nil error", plain, got, err)
		}
		if got, err := DecodeURL(EncodeURL(plain)); err != nil || got != plain {
			t.Errorf("DecodeURL(EncodeURL(%q)) = %q, %v; want the original string and a nil error", plain, got, err)
		}
	}
	if user, pass, err := ParseBasicAuth(BasicAuth("gopher", "s3cr:t")); err != nil || user != "gopher" || pass != "s3cr:t" {
		t.Errorf(`ParseBasicAuth(BasicAuth("gopher", "s3cr:t")) = %q, %q, %v; want "gopher", "s3cr:t", nil`, user, pass, err)
	}
}
