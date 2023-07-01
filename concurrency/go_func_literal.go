package main

import (
	"fmt"
	"time"
)

func Announce(message string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		fmt.Println(message)
	}() // Note the parentheses - must call the function.
}

func main() {
	Announce("Bye, world!", time.Second)

	time.Sleep(time.Duration(2) * time.Second)
}
