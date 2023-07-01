package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	// x = append(x, 4, 5, 6)
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}
