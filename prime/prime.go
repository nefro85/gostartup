package main

import (
	"gostartup/utils"
	"fmt"
)

func main()  {
	count := 0;
	var number uint64 = 0;

	for {
		if utils.IsPrime(number) {
			fmt.Printf("Prime#%d: %d\n", count + 1, number)

			count+= 1
		}
		number++

		if count >= 100 {
			break
		}
	}
}



