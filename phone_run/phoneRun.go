package main

import "gostartup/phone"

func main() {



	devices := []phone.Phone{phone.Nexus{}, phone.Iphone{}}

	for _,dev := range devices{
		dev.Call([]int{2,4,6,3})
	}





}
