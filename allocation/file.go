package main

import "fmt"

type File struct {
	fd      int
	name    string
	dirinfo *File
	nepipe  int
}

func NewFile1(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := new(File)
	f.fd = fd
	f.name = name
	f.dirinfo = nil
	f.nepipe = 0
	return f
}

func NewFile2(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	//return &File{fd, name, nil, 0}
	return &File{fd: fd, name: name}
}

func main() {
	fd := 1
	name := "data.txt"

	fmt.Printf("v1: NewFile(%d, %s) = %v\n", fd, name, NewFile1(fd, name))
	fmt.Printf("v2: NewFile(%d, %s) = %v\n", fd, name, NewFile2(fd, name))
}
