package main

import (
	"net"
	"fmt"
	"io"
)

func handleConnection(c net.Conn) {
	buff := make([]byte, 1024)

	for {
		n , err := c.Read(buff)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		s := string(buff[:n])



		fmt.Println(s)
	}
}

func main()  {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go handleConnection(conn)
	}
}




