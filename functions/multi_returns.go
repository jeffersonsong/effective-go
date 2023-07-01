package main

import "fmt"

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func main() {
	b := []byte{'1', '2', '3', 'a'}
	var x int
	for i := 0; i < len(b); {
		x, i = nextInt(b, i)
		fmt.Println(x)
	}
}
