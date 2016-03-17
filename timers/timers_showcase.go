package main

import (
	"fmt"
	"time"
	"gostartup/utils"
)

func main() {
	ch1 := make(chan string, 1)
	quit := make(chan bool, 1)

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

	fmt.Println("end begin")
	time.Sleep(time.Second * 7)
	fmt.Println("the end")
}

func asyncTaskTimeouter(quit chan bool, out chan string) {

	select {
	case <-quit:
		fmt.Println("goQuit")
		return
	default:
		utils.IsPrime(utils.LARGEST_64BIT_PRIME)

		out <- "taskdone"
	}
}
