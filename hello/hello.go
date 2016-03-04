package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("GO wrong++ ;)")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}