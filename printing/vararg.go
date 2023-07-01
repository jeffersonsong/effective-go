package main

import "fmt"

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {
	smallest := Min(5, 4, 3, 2, 1)
	fmt.Printf("min = %d\n", smallest)
}
