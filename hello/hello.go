package main

import (
	"bufio"
	"fmt"
	"os"
)

func hello() {
	fmt.Println("Hello, I'm Go!")
}

func readName() (text string) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return
}

func main() {

	hello()
	name := readName()
	fmt.Println(name)
}
