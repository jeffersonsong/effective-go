package main

import "fmt"

type ByteSlice []byte

/*
func (slice ByteSlice) Append(data []byte) []byte {
	// Body exactly the same as the Append function defined above.
}*/

func (p *ByteSlice) Append(data []byte) {
	slice := *p
	// Body as above, without the return.
	*p = append(slice, data...)
}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	slice.Append(data)
	*p = slice
	return len(data), nil
}

func main() {
	var b ByteSlice
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Println(string(b))
}
