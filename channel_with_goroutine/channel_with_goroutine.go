package main

import (
	"fmt"
	"time"
)

func main() {

	input := make(chan string, 2)
	go receive(input)

	fmt.Println("waiting for data...")
	for data := range input {
		fmt.Println(data)
	}

	fmt.Println("exit")

}

func receive(data chan<- string) {

	time.Sleep(time.Second * 5)

	data <- "data part 1"

	time.Sleep(time.Second * 5)

	data <- "data part 2"

	close(data)
}
