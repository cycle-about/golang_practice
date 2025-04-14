package main

import (
    "fmt"
    "golang.org/x/tour/pic"
    "image"
    "image/color"
)

/*
Remember the picture generator you wrote earlier? (https://go.dev/tour/moretypes/18) [my implementation: tour_3_ex_picture.go]
Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods (https://pkg.go.dev/image#Image), and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one. 
*/

type Image struct{}

func (Image) At(int, int) color.Color {
    return color.RGBA{255, 255, 255, 255}
}

func (Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, 1, 1)
}

func (Image) ColorModel() color.Model {
    return color.RGBAModel
}

func tour_4_ex_image() {
    fmt.Println("hello from image")
    m := Image{}
    pic.ShowImage(m)
}

/*
Output in both the online version and local terminal:

IMAGE:iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAEElEQVR42mL6//8/IAAA//8GBgMAX2e+rgAAAABJRU5ErkJggg==
*/