package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {

	http.HandleFunc("/hello", hndHello)
	http.ListenAndServe(":8080", nil)
}

func hndHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world" )

	fmt.Println(req)
}
