package main

import "fmt"

func main() {

	weekDay := map[int]string{
		0: "Niedziela",
		1: "Poniedziałek",
		2: "Wtorek",
		3: "Sroda",
		4: "Czwartek",
		5: "Piątek",
		6: "Sobota",
	}

	for key, value := range weekDay {
		fmt.Printf("key: %d, value : %s\n", key, value)

	}

}
