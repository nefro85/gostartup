package utils

import (
	"bytes"
	"io"
	"log"
	"net"
	"strings"
)

const BUFFER_SIZE = 32 * 1024
const LARGEST_64BIT_PRIME  = 18446744073709551557

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

func IsPrime(n uint64) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	var i uint64 = 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}
