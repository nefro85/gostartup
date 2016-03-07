package main

import (
	"fmt"
	"gostartup/command"
	"gostartup/utils"
	"net"
)

var (
	cmdIds uint64 = 100
	done          = make(chan bool, 1)
)

func main() {
	defer quit(utils.Server(handleConnection))
	<-done
}

func quit(ln net.Listener) {
	err := ln.Close()

	if err != nil {
		panic(err)
	}
	fmt.Println("server ends running")
}

func handleConnection(c net.Conn) {

	fmt.Println("handling connecton " + c.RemoteAddr().String())

	for {
		processCommand(utils.ReadCommand(c))
	}
}
func processCommand(text string) {
	var cmd command.Command

	cmd.Fill(genUid(), text)
	cmd.Execute()

	if cmd.IsQuit() {
		done <- true
	}
}

func genUid() (uid uint64) {
	// named result parameter
	uid = cmdIds
	cmdIds++
	return
}
