package main

import (
	"net/http"
	"io"
	"time"
	"fmt"
)

type httpHandler struct {

}

func (c *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world @" + time.Now().String())

	fmt.Println(r)
}

func main() {

	myHandler := &httpHandler{}

	http.Handle("/", myHandler)
	http.ListenAndServe(":8080", nil)

}


