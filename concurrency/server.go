package main

import (
	"fmt"
)

const MaxOutstanding = 5

type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

func handle(queue chan *Request) {
	for r := range queue {
		process(r)
	}
}

func process(r *Request) {
	v := r.f(r.args)
	r.resultChan <- v
}

func Serve(clientRequests chan *Request, quit chan bool) {
	// Start handlers
	for i := 0; i < MaxOutstanding; i++ {
		go handle(clientRequests)
	}
	<-quit // Wait to be told to exit.
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func main() {
	clientRequests := make(chan *Request)
	quit := make(chan bool)

	go Serve(clientRequests, quit)

	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	// Send request
	clientRequests <- request
	// Wait for response.
	quit <- true

	//time.Sleep(time.Second)
	fmt.Printf("answer: %d\n", <-request.resultChan)
}
