package main

import "fmt"

type LinesOfText [][]byte // A slice of byte slices.

func main() {
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}

	for _, line := range text {
		fmt.Println(string(line))
	}
}
