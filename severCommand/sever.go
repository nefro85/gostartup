package main

import (
	"bytes"
	"fmt"
	"gostartup/command"
	"io"
	"net"
	"strings"
)

var (
	cmdIds uint64 = 100
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
		processCommand(readCommand(c))
	}
}
func processCommand(text string) {
	var cmd command.Command

	cmd.Fill(genUid(), text)
	cmd.Execute()

}

func genUid() (uid uint64) { // named result parameter
	uid = cmdIds
	cmdIds++
	return
}


func readCommand(c net.Conn) string {
	var cmdBuff bytes.Buffer
	buff := make([]byte, 32*1024)
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

