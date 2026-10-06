//go:build solution

package regularexpressions

import "regexp"

var (
	dateRe    = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	hashtagRe = regexp.MustCompile(`#\w+`)
	digitRe   = regexp.MustCompile(`\d`)
	varRe     = regexp.MustCompile(`\$\{\w+\}`)
)

func IsHexColor(s string) bool {
	match, _ := regexp.MatchString(`^#[0-9a-fA-F]{6}$`, s)
	return match
}

func FindFirst(pattern, text string) (string, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	return r.FindString(text), nil
}

func ParseDate(s string) (year, month, day string, ok bool) {
	m := dateRe.FindStringSubmatch(s)
	if m == nil {
		return "", "", "", false
	}
	// m[0] is the whole match, the groups start at index 1.
	return m[1], m[2], m[3], true
}

func Hashtags(text string) []string {
	return hashtagRe.FindAllString(text, -1)
}

func MaskDigits(s string) string {
	return digitRe.ReplaceAllString(s, "#")
}

func ExpandVars(text string, vars map[string]string) string {
	return varRe.ReplaceAllStringFunc(text, func(match string) string {
		name := match[2 : len(match)-1] // strip "${" and "}"
		if value, ok := vars[name]; ok {
			return value
		}
		return match
	})
}
