package main

import (
	"bufio"
	"fmt"
	"os"
)

func p1() {
	fmt.Println("GO wrong++ ;)")
}

func main() {

	p1()
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
