package utils

import (
	"bytes"
	"io"
	"log"
	"net"
	"strings"
)

const BUFFER_SIZE = 32 * 1024

func ReadCommand(c net.Conn) string {
	var cmdBuff bytes.Buffer
	buff := make([]byte, BUFFER_SIZE)
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

func Server(hnd Serve) (ln net.Listener) {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("can't listen")
	}

	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Print("can't accept")
				break
			} else {
				go hnd(conn)
			}
		}

	}()

	return

}
