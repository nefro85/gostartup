package main

import (
	"fmt"
	pn "gostartup/utils"
)

func main() {
	//trivial()
	byChannel()

}

func trivial() {
	count := 0
	var number uint64 = 0

	for {
		if pn.IsPrime(number) {
			fmt.Printf("Prime#%d: %d\n", count+1, number)

			count += 1
		}
		number++

		if count >= 100 {
			break
		}
	}
}

func byChannel() {

	channel := make(chan uint64)

	go func() {
		from := pn.STRESS_PRIMES
		quantity := 10
		pn.GenPrime(from, quantity, channel)
		close(channel)
	}()

	for prime := range channel {
		fmt.Println(prime)
	}
}
