package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/hello", hndHello)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func hndHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world @"+time.Now().String())

	fmt.Println(req)
}
