// this is the version that runs on the Go tour sandbox
// DOES NOT show an image when run in terminal
// when run in sandbox, it RUNS IT IN A TEST MODE THAT PROVIDES INPUTS
// exercise 3.23 on maps actually shows in main "wc.Test(WordCount)"

package main

// import "golang.org/x/tour/pic"

// func Pic(h, w int) [][]uint8 {
// 	picture := [][]uint8{}		// this way makes nil structure
// 	for i:=0; i < h; i++ {
// 		rowSlice := populate_row(w, i)
// 		picture = append(picture, rowSlice)
// 	}
// 	return picture
// }

// func populate_row(w, r int) []uint8 {
// 	rowSlice := make([]uint8, w)
// 	for i:= 0; i < w; i++ {
// 		// rowSlice[i] = uint8((i+r)/2)
// 		rowSlice[i] = uint8(i*r)
// 		// rowSlice[i] = uint8(i^r)
// 	}
// 	return rowSlice
// }

// func main() {
// 	pic.Show(Pic)
// }