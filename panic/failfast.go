package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
	fmt.Printf("USER: %v\n", user)
}

func main() {

}
