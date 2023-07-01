package main

import (
	"log"
	"time"
)

type Work int

func server(workChan <-chan *Work) {
	for work := range workChan {
		go safelyDo(work)
	}
}

func safelyDo(work *Work) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(work)
}

func do(work *Work) {
	if int(*work)%2 == 0 {
		panic("Some bad thing happen!")
	}
}

func main() {
	workChan := make(chan *Work)
	go server(workChan)

	for i := 0; i < 10; i++ {
		work := Work(i)
		workChan <- &work
	}

	time.Sleep(time.Second)
}
