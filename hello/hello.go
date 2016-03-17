package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fn func(func())

var name string

func main() {

	enterYourName := func() {
		name = readName()

	}
	hello(whatsIsYourName, enterYourName)

}

func hello(ask fn, readName func()) {
	defer func() { fmt.Println(name + ", good luck and have fun!") }()
	fmt.Println("Hello, I'm Go!")
	ask(readName)
}

func whatsIsYourName(read func()) {
	fmt.Print("What is your name..?")
	read()
}

func readName() (text string) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	if err != nil {
		fmt.Println(err)
	}
	return
}
