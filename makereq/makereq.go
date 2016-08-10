package main

import (
	"flag"
	"fmt"
	"net/http"
)

func worker(path string, requests <-chan int, result chan<- int) {
	for j := range requests {
		fmt.Println(j)
		_, err := http.Get(fmt.Sprintf("%s/kafka_tracking/%d", path, j))
		if err != nil {
			panic(err)
		}

		result <- j
	}
}

func main() {
	numRequests := flag.Int("messages", 100, "number of messages")
	numThreads := flag.Int("threads", 20, "number of threads")
	path := flag.String("path", "http://localhost:3000", "path to request")
	requests := make(chan int, *numThreads)
	results := make(chan int, *numRequests)

	for w := 1; w <= *numThreads; w++ {
		go worker(*path, requests, results)
	}

	for j := 1; j <= *numRequests; j++ {
		requests <- j
	}
	close(requests)

	for i := 0; i < *numRequests; i++ {
		fmt.Println(<-results)
	}
}
