package main

import "fmt"

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func main() {
	fmt.Printf("unhex(%c) = %d\n", '5', unhex('5'))
	fmt.Printf("unhex(%c) = %d\n", 'b', unhex('b'))
	fmt.Printf("unhex(%c) = %d\n", 'B', unhex('B'))
}
