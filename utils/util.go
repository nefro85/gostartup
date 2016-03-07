package utils

import (
	"net"
	"bytes"
	"io"
	"strings"
)
func ReadCommand(c net.Conn) string {
	var cmdBuff bytes.Buffer
	buff := make([]byte, 32 * 1024)
	for {
		n, err := c.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		s := string(buff[:n])

		cmdBuff.WriteString(s)
		if strings.Contains(s, "\r\n") {
			break
		}
	}
	return cmdBuff.String()
}

type Serve func(net.Conn)

func Server(hnd Serve)  {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("can't listen")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("can't accept")
		}
		go hnd(conn)
	}

}


