package utils

import (
	"gocv.io/x/gocv"
	"image"
)

func IMRead720p(filename string) gocv.Mat {
	img := gocv.IMRead(filename, gocv.IMReadColor)
	img720p := gocv.NewMatWithSize(1280, 720, gocv.MatTypeCV8UC3)
	gocv.Resize(img, &img720p, image.Point{1280, 720}, 0.0, 0.0, gocv.InterpolationLinear)
	return img720p
}
