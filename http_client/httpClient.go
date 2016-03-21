package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UrlReq struct {
	Url      string
	Response string
}

const (
	BUFF = 10
)

func main() {
	jobs := make(chan UrlReq, BUFF)
	responses := make(chan UrlReq, BUFF)

	go performHttpRequest(jobs, responses)

	urls := make([]UrlReq, BUFF)
	for i := 0; i < BUFF; i++ {
		urls[i] = UrlReq{"http://localhost:8080/hello?uid=" + strconv.Itoa(i + 1), ""}
		//log.Println(urls[i])
	}

	for _, j := range urls {
		jobs <- j
	}
	close(jobs)

	for out := range responses {
		fmt.Println(out)
	}
}

func performHttpRequest(requests <-chan UrlReq, output chan<- UrlReq) {
	for r := range requests {
		go func(item UrlReq) {
			resp, _ := http.Get(item.Url)
			data, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				log.Fatal(err)
			} else {

				item.Response = string(data[:len(data)])
				output <- item
			}
		}(r)
	}
}
