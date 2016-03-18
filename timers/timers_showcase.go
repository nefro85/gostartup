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

	time.Sleep(time.Second * 10)
	console.Println("exit")
	os.Exit(0)
}

func application() {
	ch1 := make(chan string, 1)
	quit := make(chan bool, 1)

	//defer func() {
	//	console.Println("clean up!")
	//	close(ch1)
	//	close(quit)
	//}()

	go asyncTask(quit, ch1)

	select {
	case data := <-ch1:
		console.Println(data)
		break
	case <-time.After(time.Second * 5):
		console.Println("timeout")
		quit <- true
		close(quit)
		close(ch1)
	}
}

func asyncTask(quit chan bool, out chan string) {

	check := func() {
		primeFmt := func(base int) string {
			return strconv.FormatUint(utils.BIG_PRIME, base)
		}

		console.Printf("async: checking is %s (dec:%s) is a prime number...\n", primeFmt(16), primeFmt(10))
		isPrime := utils.IsPrime(utils.BIG_PRIME)

		out <- "async: taskdone, is prime: " + strconv.FormatBool(isPrime)
		//primeResult <- isPrime
	}

	select {
	case <-quit:
		console.Println("async: quit async task")
		return
	//case <-time.After(time.Second * 3):
	//	console.Println("async: quit timeout")
	//	return
	default:
		check()
	}
}
