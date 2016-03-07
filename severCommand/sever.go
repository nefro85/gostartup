package main

import (
	"fmt"
	"gostartup/command"
	"gostartup/utils"
	"net"
)

var (
	cmdIds uint64 = 100
)

func main() {
	utils.Server(handleConnection)
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

}

func genUid() (uid uint64) {
	// named result parameter
	uid = cmdIds
	cmdIds++
	return
}


