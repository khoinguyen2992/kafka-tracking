package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func worker(path string, requests <-chan int, result chan<- int) {
	for j := range requests {
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
	flag.Parse()
	requests := make(chan int, *numRequests)
	results := make(chan int, *numRequests)
	completed := make(chan int, *numThreads)
	fmt.Printf("Num of messages: %d\n", *numRequests)
	fmt.Printf("Num of threads: %d\n", *numThreads)

	fmt.Printf("Workers creating at: %s\n", time.Now())
	for w := 1; w <= *numThreads; w++ {
		go func(path string, requests <-chan int, result chan<- int, completed chan<- int) {
			completed <- w
			worker(path, requests, results)
		}(*path, requests, results, completed)
	}
	for i := 1; i <= *numThreads; i++ {
		<-completed
	}
	fmt.Printf("Workers created at: %s\n", time.Now())

	fmt.Printf("Jobs creating at: %s\n", time.Now())
	for j := 1; j <= *numRequests; j++ {
		requests <- j
	}
	close(requests)
	fmt.Printf("Jobs created at: %s\n", time.Now())

	for i := 1; i <= *numRequests; i++ {
		fmt.Println(<-results)
	}
}
