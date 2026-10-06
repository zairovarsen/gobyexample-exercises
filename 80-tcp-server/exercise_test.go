package tcpserver

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

// Every network operation in these tests has a deadline, so a server that
// does not answer makes a test fail after two seconds instead of hanging.
const timeout = 2 * time.Second

// client is the test's end of a connection.
type client struct {
	t    *testing.T
	conn net.Conn
	r    *bufio.Reader
}

func newClient(t *testing.T, conn net.Conn) *client {
	t.Helper()
	t.Cleanup(func() { conn.Close() })
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		t.Fatal(err)
	}
	return &client{t, conn, bufio.NewReader(conn)}
}

// send writes text to the server.
func (c *client) send(text string) {
	c.t.Helper()
	if _, err := c.conn.Write([]byte(text)); err != nil {
		c.t.Fatalf("sending %q failed: %v (is the server reading from the connection?)", text, err)
	}
}

// expect reads one line from the server and compares it.
func (c *client) expect(want string) {
	c.t.Helper()
	got, err := c.r.ReadString('\n')
	if err != nil {
		c.t.Fatalf("waiting for the response %q: read failed with %v after receiving %q", want, err, got)
	}
	if got != want {
		c.t.Errorf("the server responded %q, want %q", got, want)
	}
}

// expectClosed checks that the server closed the connection: reading then
// reports io.EOF.
func (c *client) expectClosed() {
	c.t.Helper()
	if b, err := c.r.ReadByte(); err != io.EOF {
		c.t.Errorf("the server did not close the connection (read returned %q, %v; want io.EOF)", b, err)
	}
}

// pipe returns the client end of an in-memory connection (net.Pipe) whose
// server end is being served by handle, the server end itself, and a channel
// that is closed when handle returns.
func pipe(t *testing.T, handle func(net.Conn)) (*client, net.Conn, <-chan struct{}) {
	t.Helper()
	clientEnd, serverEnd := net.Pipe()
	t.Cleanup(func() { serverEnd.Close() })
	done := make(chan struct{})
	go func() {
		defer close(done)
		handle(serverEnd)
	}()
	return newClient(t, clientEnd), serverEnd, done
}

// expectHangUpHandled closes the client's end of a pipe and checks that the
// handler notices: it has to return and to close its own end.
func expectHangUpHandled(t *testing.T, name string, c *client, serverEnd net.Conn, done <-chan struct{}) {
	t.Helper()
	c.conn.Close()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Fatalf("%s did not return after the client closed the connection", name)
	}
	// Reading from a pipe end that was closed locally reports
	// io.ErrClosedPipe; io.EOF would mean only the client's end is closed.
	if _, err := serverEnd.Read(make([]byte, 1)); !errors.Is(err, io.ErrClosedPipe) {
		t.Errorf("%s left its connection open after the client disconnected (read error = %v); use defer conn.Close()", name, err)
	}
}

// listen opens a TCP listener on a free port of this machine.
func listen(t *testing.T) net.Listener {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("cannot listen on a local port: %v", err)
	}
	t.Cleanup(func() { ln.Close() })
	return ln
}

// dial connects to a listener created by listen.
func dial(t *testing.T, ln net.Listener) *client {
	t.Helper()
	conn, err := net.DialTimeout("tcp", ln.Addr().String(), timeout)
	if err != nil {
		t.Fatalf("cannot connect to the server: %v", err)
	}
	return newClient(t, conn)
}

// startServe runs Serve(ln, handle) in a goroutine and returns the channel
// that delivers its result. It fails the test right away when Serve returns
// immediately instead of waiting for connections.
func startServe(t *testing.T, ln net.Listener, handle func(net.Conn)) <-chan error {
	t.Helper()
	served := make(chan error, 1)
	go func() { served <- Serve(ln, handle) }()
	select {
	case err := <-served:
		t.Fatalf("Serve returned %v right away; it must keep accepting connections until the listener is closed", err)
	case <-time.After(50 * time.Millisecond):
	}
	return served
}

// stopServe closes the listener and checks that Serve then returns nil.
func stopServe(t *testing.T, ln net.Listener, served <-chan error) {
	t.Helper()
	ln.Close()
	select {
	case err := <-served:
		if err != nil {
			t.Errorf("Serve returned %v after the listener was closed, want nil (check errors.Is(err, net.ErrClosed))", err)
		}
	case <-time.After(timeout):
		t.Error("Serve did not return after the listener was closed")
	}
}

func TestAck(t *testing.T) {
	tests := []struct {
		message string
		want    string
	}{
		{"Hello TCP Server\n", "ACK: HELLO TCP SERVER\n"},
		{"hello\r\n", "ACK: HELLO\n"},
		{"  padded \t\n", "ACK: PADDED\n"},
		{"no newline", "ACK: NO NEWLINE\n"},
		{"\n", "ACK: \n"},
	}
	for _, tt := range tests {
		if got := Ack(tt.message); got != tt.want {
			t.Errorf("Ack(%q) = %q, want %q", tt.message, got, tt.want)
		}
	}
}

func TestHandleConn(t *testing.T) {
	c, _, done := pipe(t, HandleConn)
	c.send("Hello TCP Server\n")
	c.expect("ACK: HELLO TCP SERVER\n")
	c.expectClosed()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("HandleConn did not return after answering")
	}

	// Windows-style line ending.
	c, _, _ = pipe(t, HandleConn)
	c.send("ping\r\n")
	c.expect("ACK: PING\n")
	c.expectClosed()

	// A client that disconnects without sending a line.
	c, serverEnd, done := pipe(t, HandleConn)
	expectHangUpHandled(t, "HandleConn", c, serverEnd, done)
}

func TestServe(t *testing.T) {
	ln := listen(t)
	// The test's own handler: answer one line with "handled <line>".
	handle := func(conn net.Conn) {
		defer conn.Close()
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		line, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			return
		}
		fmt.Fprintf(conn, "handled %s", line)
	}
	served := startServe(t, ln, handle)

	// The first client connects but stays silent for now. Its handler is
	// stuck waiting for a line.
	slow := dial(t, ln)

	// A second client must be served meanwhile. That only works when every
	// connection is handled in its own goroutine.
	quick := dial(t, ln)
	quick.send("quick\n")
	if got, err := quick.r.ReadString('\n'); err != nil {
		t.Fatalf("a second client got no answer while the first connection was still open (%v); "+
			"start a goroutine per connection: go handle(conn)", err)
	} else if got != "handled quick\n" {
		t.Errorf("the second client received %q, want %q", got, "handled quick\n")
	}

	// The first client is still connected and gets its answer too.
	slow.send("slow\n")
	slow.expect("handled slow\n")

	// And the loop keeps accepting.
	third := dial(t, ln)
	third.send("third\n")
	third.expect("handled third\n")

	stopServe(t, ln, served)
}

// failingListener is a listener whose Accept always fails.
type failingListener struct{ err error }

func (l failingListener) Accept() (net.Conn, error) { return nil, l.err }
func (l failingListener) Close() error              { return nil }
func (l failingListener) Addr() net.Addr            { return &net.TCPAddr{} }

func TestServeAcceptError(t *testing.T) {
	errAccept := errors.New("accept failed")
	served := make(chan error, 1)
	go func() {
		served <- Serve(failingListener{errAccept}, func(conn net.Conn) { conn.Close() })
	}()
	select {
	case err := <-served:
		if !errors.Is(err, errAccept) {
			t.Errorf("Serve returned %v when Accept failed, want the error from Accept", err)
		}
	case <-time.After(timeout):
		t.Fatal("Serve did not return although Accept failed; return errors other than net.ErrClosed")
	}
}

func TestServeWithHandleConn(t *testing.T) {
	// The complete server from the page.
	ln := listen(t)
	served := startServe(t, ln, HandleConn)

	// Five clients are connected at the same time ...
	clients := make([]*client, 5)
	for i := range clients {
		clients[i] = dial(t, ln)
	}
	// ... and are answered in the opposite order of their arrival.
	for i := len(clients) - 1; i >= 0; i-- {
		clients[i].send(fmt.Sprintf("hello from client %d\n", i))
		clients[i].expect(fmt.Sprintf("ACK: HELLO FROM CLIENT %d\n", i))
	}

	stopServe(t, ln, served)
}

// testServer starts a minimal TCP server for the Send tests (independent of
// the Serve exercise) and returns its address. handle runs once for every
// connection; the connection is closed when handle returns.
func testServer(t *testing.T, handle func(conn net.Conn)) string {
	t.Helper()
	ln := listen(t)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			go func() {
				defer conn.Close()
				conn.SetDeadline(time.Now().Add(5 * time.Second))
				handle(conn)
			}()
		}
	}()
	return ln.Addr().String()
}

type sendResult struct {
	response string
	err      error
}

// send calls Send in a goroutine and fails the test instead of hanging when
// Send does not return in time.
func send(t *testing.T, addr, message string, limit time.Duration) sendResult {
	t.Helper()
	done := make(chan sendResult, 1)
	go func() {
		response, err := Send(addr, message, limit)
		done <- sendResult{response, err}
	}()
	select {
	case res := <-done:
		return res
	case <-time.After(5 * time.Second):
		t.Fatalf("Send(addr, %q, %v) did not return - set a deadline on the connection", message, limit)
	}
	panic("unreachable")
}

func TestSend(t *testing.T) {
	received := make(chan string, 1)
	closed := make(chan error, 1)
	addr := testServer(t, func(conn net.Conn) {
		r := bufio.NewReader(conn)
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		received <- line
		fmt.Fprintf(conn, "ECHO: %s", strings.ToUpper(line))
		// After the answer the client is expected to hang up.
		_, err = r.ReadByte()
		closed <- err
	})

	res := send(t, addr, "hello", timeout)
	if res.err != nil {
		t.Fatalf(`Send(addr, "hello", 2s) returned error: %v (the message must end in "\n" for the server to answer)`, res.err)
	}
	if want := "ECHO: HELLO"; res.response != want {
		t.Errorf(`Send(addr, "hello", 2s) = %q, want %q (without the line ending)`, res.response, want)
	}
	select {
	case line := <-received:
		if line != "hello\n" {
			t.Errorf(`the server received %q, want "hello\n"`, line)
		}
	default:
		t.Fatal("the server never received a line from Send")
	}
	select {
	case err := <-closed:
		if err != io.EOF {
			t.Errorf("after Send returned, the server's read ended with %v, want io.EOF from a closed connection", err)
		}
	case <-time.After(timeout):
		t.Error("Send did not close the connection (defer conn.Close())")
	}
}

func TestSendTimeout(t *testing.T) {
	// This server accepts the connection and then never answers.
	stop := make(chan struct{})
	defer close(stop)
	silent := testServer(t, func(conn net.Conn) {
		<-stop
	})

	start := time.Now()
	res := send(t, silent, "anyone there?", 100*time.Millisecond)
	if res.err == nil {
		t.Errorf("Send(silent server, 100ms) = %q, nil - want a timeout error", res.response)
	}
	if elapsed := time.Since(start); elapsed > 1500*time.Millisecond {
		t.Errorf("Send(silent server, 100ms) took %v, want it to give up after about 100ms", elapsed)
	}

	// Nobody listens on this address any more.
	ln := listen(t)
	gone := ln.Addr().String()
	ln.Close()
	if res := send(t, gone, "hello", 200*time.Millisecond); res.err == nil {
		t.Errorf("Send(address without a server) = %q, nil - want the error from net.DialTimeout", res.response)
	}
}

func TestHandleSession(t *testing.T) {
	c, _, done := pipe(t, HandleSession)
	c.send("hello\n")
	c.expect("ACK: HELLO\n")
	c.send("  go  \r\n")
	c.expect("ACK: GO\n")
	c.send("quit it\n") // not the quit command
	c.expect("ACK: QUIT IT\n")
	c.send("Quit\n")
	c.expect("BYE\n")
	c.expectClosed()
	select {
	case <-done:
	case <-time.After(timeout):
		t.Error("HandleSession did not return after the client said quit")
	}

	// A client that disconnects without saying quit.
	c, serverEnd, done := pipe(t, HandleSession)
	c.send("one\n")
	c.expect("ACK: ONE\n")
	expectHangUpHandled(t, "HandleSession", c, serverEnd, done)
}
