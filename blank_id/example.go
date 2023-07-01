package main

import (
	"log"
	"os"

	_ "net/http/pprof"
)

func main() {
	fd, err := os.Open("test.go")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: use fd.
	_ = fd
}
