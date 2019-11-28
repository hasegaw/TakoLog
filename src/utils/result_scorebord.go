package utils

import (
	"gocv.io/x/gocv"
	"image"
)

func IMRead720p(filename string) gocv.Mat {
	img := gocv.IMRead(filename, gocv.IMReadColor)
	img_720p := gocv.NewMatWithSize(1280, 720, gocv.MatTypeCV8UC3)
	gocv.Resize(img, &img_720p, image.Point{1280, 720}, 0.0, 0.0, gocv.InterpolationLinear)
	return img_720p
}
