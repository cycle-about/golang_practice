// drafts commented out to avoid name conflicts with sandbox version

package main

import (
	// "golang.org/x/tour/pic"
)

	/* 
	Tour part 3, page 18 Exercies: Slices
	Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. 
	When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.
	The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.
	(You need to use a loop to allocate each []uint8 inside the [][]uint8.)
	(Use uint8(intValue) to convert between types.) 
	*/

func Run_tour_3_ex() {
	// Pic(3, 1)
	// pic.Show(Pic)

	// Pic(10, 4)
	// pic.Show(Pic)
}

// if not passed dx and dy are 256
// func Pic(w, h int) [][]uint8 {
// 	return make_2D_array(w, h)
// }

// row=0 len=0 cap=0 []
// row=1 len=0 cap=0 []
// row=2 len=0 cap=0 []
// func make_2D_array_empty_rows(dx, dy int) {
// 	picture := make([][]uint8, dy) 		// creates dx slices each with len and cap 0
// 	fmt.Println("2D array with make")
// 	print2DSlice(picture)
// }

// row=0 len=4 cap=4 [0 0 0 0]
// row=1 len=4 cap=4 [0 0 0 0]
// row=2 len=4 cap=4 [0 0 0 0]
// func make_2D_array(w, h int) [][]uint8 {
// 	picture := [][]uint8{}		// this way makes nil structure
// 	for i:=0; i < h; i++ {
// 		rowSlice := populate_row(w, i)
// 		picture = append(picture, rowSlice)
// 	}
// 	print2DSliceShort(picture)
// 	return picture
// }

// func populate_row(w, r int) []uint8 {
// 	rowSlice := make([]uint8, w)
// 	for i:= 0; i < w; i++ {
// 		// rowSlice[i] = uint8((i+r)/2)
// 		// rowSlice[i] = uint8(i*r)
// 		rowSlice[i] = uint8(i^r)
// 	}
// 	return rowSlice
// }