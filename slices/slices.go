package main

import "fmt"

func main() {

	p := fmt.Println

	d := make([]string, 2)

	p("cap: ", cap(d))
	p("len: ", len(d))

	d[0] = "N"
	d[1] = "O"
	p(d)

	d = append(d, "K")
	d = append(d, "I", "A")

	p(d)

	p("cap: ", cap(d))
	p("len: ", len(d))

	p(d[0:2])
	p(d[2:5])

	p(d[0:])

}
