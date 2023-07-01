package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(http.ResponseWriter, *http.Request)

// ServeHTTP calls f(w, req).
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	f(w, req)
}

// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	flag.Parse()
	ctr := new(Counter)
	http.Handle("/counter", ctr)
	http.Handle("/args", http.HandlerFunc(ArgServer))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
