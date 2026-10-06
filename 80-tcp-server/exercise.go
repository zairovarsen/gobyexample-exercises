//go:build !solution

// Package tcpserver contains exercises for https://gobyexample.com/tcp-server
//
// Run the tests for this section with:
//
//	go test ./80-tcp-server/
package tcpserver

import (
	"net"
	"time"
)

// The server on the page listens on port 8090, accepts connections in an
// endless loop and handles every connection in its own goroutine: it reads
// one line, answers "ACK: " plus the line in upper case, and closes the
// connection. The exercises build that server from testable pieces. Nothing
// listens on a fixed port: the tests create the listener on a free port and
// pass it in.

// Ack builds the server's answer to one received message: surrounding white
// space (including the line ending "\n" or "\r\n") is trimmed, the rest is
// upper-cased and wrapped as "ACK: <MESSAGE>\n".
//
// Example: Ack("Hello TCP Server\n") is "ACK: HELLO TCP SERVER\n".
func Ack(message string) string {
	// TODO: trim, upper-case, format.
	return ""
}

// HandleConn serves one client connection, like handleConnection on the
// page: it reads a single line (everything up to and including "\n") from
// conn, writes Ack(line) back, and closes the connection.
//
// The connection must be closed in every case - also when reading fails
// because the client hung up without sending a line; nothing is written
// then. "defer conn.Close()" as the first statement takes care of it.
//
// Hint: bufio.NewReader(conn).ReadString('\n').
func HandleConn(conn net.Conn) {
	// TODO: defer the close, read a line, write the answer.
}

// Serve is the accept loop of the server: it accepts connections from ln
// again and again and starts handle(conn) in a NEW GOROUTINE for each one,
// so that a slow client does not keep the others waiting. On the page this
// is the for loop around listener.Accept() with "go handleConnection(conn)".
//
// Serve only returns when Accept fails. There are two cases:
//   - the listener was closed (that is how a server is shut down):
//     errors.Is(err, net.ErrClosed) is true and Serve returns nil;
//   - any other error is returned as it is. (The page logs the error and
//     carries on; returning it keeps the exercise simple.)
//
// A real program would call:
//
//	ln, err := net.Listen("tcp", ":8090")
//	...
//	err = Serve(ln, HandleConn)
func Serve(ln net.Listener, handle func(net.Conn)) error {
	// TODO: loop: accept, check the error, go handle(conn).
	return nil
}

// Send is the matching client (the page uses the nc tool instead): it
// connects to the TCP server at addr (e.g. "127.0.0.1:8090"), sends message
// followed by "\n", reads one line of response and returns it without the
// line ending. The connection is closed before Send returns.
//
// Example: against the page's server, Send(addr, "hello", time.Second)
// returns "ACK: HELLO".
//
// Nothing may block for longer than timeout: connect with
// net.DialTimeout("tcp", addr, timeout), and limit the reading and writing
// with conn.SetDeadline(time.Now().Add(timeout)). When the server cannot be
// reached or does not answer in time, Send returns "" and the error.
func Send(addr, message string, timeout time.Duration) (string, error) {
	// TODO: dial, defer the close, set the deadline, write, read a line.
	return "", nil
}

// HandleSession is a talkative variant of HandleConn that keeps the
// connection open for a whole conversation. It reads line after line and
// answers each one with Ack(line) - until the client says "quit" (in any
// mix of upper and lower case, surrounding white space ignored). The answer
// to quit is "BYE\n", after which the connection is closed.
//
// The session also ends, without a goodbye, when the client closes the
// connection. In every case conn must be closed when HandleSession returns.
//
// Example conversation (-> received, <- sent):
//
//	-> hello
//	<- ACK: HELLO
//	-> go
//	<- ACK: GO
//	-> quit
//	<- BYE
//
// Hint: a bufio.Scanner over conn delivers the lines one by one, exactly
// like the line filters did for os.Stdin; strings.EqualFold compares
// ignoring case.
func HandleSession(conn net.Conn) {
	// TODO: defer the close, scan lines, answer each, stop at "quit".
}
