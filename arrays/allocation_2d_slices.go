package main

import "fmt"

func allocate1(XSize, YSize int) [][]uint8 {
	// Allocate the top-level slice.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Loop over the rows, allocating the slice for each row.
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}

	return picture
}

func allocate2(XSize, YSize int) [][]uint8 {
	// Allocate the top-level slice, the same as before.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture[i], pixels = pixels[:XSize], pixels[XSize:]
	}

	return picture
}

func dimension(matrix [][]uint8) (x, y int) {
	b := cap(matrix)
	a := cap(matrix[0])

	return a, b
}

func main() {
	matrix1 := allocate1(10, 5)
	x1, y1 := dimension(matrix1)
	fmt.Printf("dim(matrix1) = (%d, %d)\n", x1, y1)

	matrix2 := allocate2(10, 5)
	x2, y2 := dimension(matrix2)
	fmt.Printf("dim(matrix2) = (%d, %d)\n", x2, y2)
}
