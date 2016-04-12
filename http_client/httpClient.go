package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UrlReq struct {
	Url      string
	Response string
}

const (
	BUFF = 100
)

func main() {
	jobs := make(chan UrlReq, BUFF)
	responses := make(chan UrlReq, BUFF)

	counter := make(chan bool, BUFF)
	go watcher(counter, responses)

	go performHttpRequest(jobs, responses, counter)

	urls := make([]UrlReq, BUFF)
	for i := 0; i < BUFF; i++ {
		urls[i] = UrlReq{"http://localhost:8080/hello?uid=" + strconv.Itoa(i+1), ""}
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

func watcher(counter chan bool, responses chan UrlReq) {
	clean := func() {
		close(responses)
		close(counter)
	}
	doCount := func() bool {
		count := 0
		for range counter {
			count++
			if count >= BUFF {
				return true
			}
		}
		return false
	}

	select {
	case <-time.After(time.Second + 5):
		log.Println("timeout")
		clean()
		break
	default:
		if doCount() {
			log.Println("http job done")
			clean()
		} else {
			log.Panic("!!!")
		}

	}

}

func performHttpRequest(requests <-chan UrlReq, output chan<- UrlReq, counter chan<- bool) {
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
			counter <- true
		}(r)
	}
}
