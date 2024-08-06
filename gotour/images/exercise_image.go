package images

/*

Exercise: Images
================

Remember the picture generator you wrote earlier? Let's write another one,
but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds
to color.RGBA{v, v, 255, 255} in this one.

*/

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.h, img.w)
}

func (img Image) At(x, y int) color.Color {
	v1 := uint8((x + y) / 2)
	v2 := uint8(math.Sqrt(float64(x*x + y*y)))
	return color.RGBA{uint8(x), uint8(y), v1, v2}
}
func ExerciseImage() {
	m := Image{200, 200}
	pic.ShowImage(m)
}
