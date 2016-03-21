package main

import (
	console "fmt"
	"gostartup/utils"
	"os"
	"strconv"
	"time"
)

func main() {

	application()

	console.Println("exit")
	os.Exit(0)
}

func application() {
	ch1 := make(chan string, 1)
	quit := make(chan bool, 1)

	defer func() {
		console.Println("clean up!")
		close(ch1)
		close(quit)
	}()

	go asyncTask(quit, ch1)

	select {
	case data := <-ch1:
		console.Println(data)
		break
	case <-time.After(time.Second * 5):
		console.Println("timeout")
		quit <- true
	}
}

func asyncTask(quit <-chan bool, out chan<- string) {
	select {
	case <-quit:
		close(out)
		console.Println("async: quit async task")
		return
	default:
		checkIsPrime(out, utils.BIG_PRIME)
	}
}

func checkIsPrime(out chan<- string, number uint64) {
	primeFmt := func(base int) string {
		return strconv.FormatUint(number, base)
	}

	console.Printf("async: checking is %s (dec:%s) is a prime number...\n", primeFmt(16), primeFmt(10))
	isPrime := utils.IsPrime(number)

	out <- "async: taskdone, is prime: " + strconv.FormatBool(isPrime)
}
