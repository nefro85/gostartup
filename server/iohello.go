package main

import (
	"fmt"
	"gostartup/utils"
	"net"
)

func main() {

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic("can't listen")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("can't accept")
		}
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {

	fmt.Println("handling connecton " + c.RemoteAddr().String())

	for {
		processCommand(utils.ReadCommand(c))
	}
}

func processCommand(cmd string) {
	fmt.Print(cmd)
}
