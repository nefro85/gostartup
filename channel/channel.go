package main

import "fmt"

func main() {

	channel := make(chan string, 2)


	channel <- "data part 1"
	channel <- "data part 2"

	close(channel)


	for data := range channel {
		fmt.Println(data)
	}
}
