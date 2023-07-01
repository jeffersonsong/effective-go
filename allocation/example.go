package main

import "fmt"

func main() {
	var p *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful
	//var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	//var p *[]int = new([]int)
	*p = make([]int, 100, 100)

	// Idiomatic:
	v := make([]int, 100)

	fmt.Println("p:", p)
	fmt.Println("v:", v)
}
