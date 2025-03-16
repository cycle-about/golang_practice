// this is the version that runs on the Go tour sandbox
// DOES NOT show an image when run in terminal

package main

import "golang.org/x/tour/pic"

func Pic(h, w int) [][]uint8 {
	picture := [][]uint8{}		// this way makes nil structure
	for i:=0; i < h; i++ {
		rowSlice := populate_row(w, i)
		picture = append(picture, rowSlice)
	}
	return picture
}

func populate_row(w, r int) []uint8 {
	rowSlice := make([]uint8, w)
	for i:= 0; i < w; i++ {
		// rowSlice[i] = uint8((i+r)/2)
		rowSlice[i] = uint8(i*r)
		// rowSlice[i] = uint8(i^r)
	}
	return rowSlice
}

func Run_tour_3_sandbox() {
	pic.Show(Pic)
}