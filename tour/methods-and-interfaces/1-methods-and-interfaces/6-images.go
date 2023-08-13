package main

import (
	"fmt"
	"image"
)

// type Image struct{}
// import "golang.org/x/tour/pic"
func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

	// Exercise: Images
	//m2 := Image{}
	//pic.ShowImage(m)
}
