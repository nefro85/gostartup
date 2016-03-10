package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	quit := make(chan bool)

	go asyncTaskTimeouter(quit, ch1)

	select {
	case data := <-ch1:
		fmt.Println(data)
		break
	case <-time.After(time.Second * 2):
		fmt.Println("timeout")
		quit <- true
		close(ch1)
	}

	time.Sleep(time.Second * 7)
	fmt.Println("the end")
}

func asyncTaskTimeouter(quit chan bool, out chan string) {

	defer fmt.Println("die")
	select {
	case <-quit:
		fmt.Println("goQuit")
		return
	default:
		time.Sleep(time.Second * 3)

		out <- "taskdone"
	}
}
