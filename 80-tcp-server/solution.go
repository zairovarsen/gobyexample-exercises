//go:build solution

package tcpserver

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"
)

func Ack(message string) string {
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	return fmt.Sprintf("ACK: %s\n", ackMsg)
}

func HandleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	conn.Write([]byte(Ack(message)))
}

func Serve(ln net.Listener, handle func(net.Conn)) error {
	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		go handle(conn)
	}
}

func Send(addr, message string, timeout time.Duration) (string, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		return "", err
	}
	if _, err := fmt.Fprintf(conn, "%s\n", message); err != nil {
		return "", err
	}
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimRight(response, "\r\n"), nil
}

func HandleSession(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.EqualFold(line, "quit") {
			conn.Write([]byte("BYE\n"))
			return
		}
		if _, err := conn.Write([]byte(Ack(line))); err != nil {
			return
		}
	}
}
