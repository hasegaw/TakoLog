package scoreboard

import (
	// "fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func ExtractFeature(img gocv.Mat) gocv.Mat {
	imgWinBGR := img.Region(image.Rect(720, 80, 720+80, 80+32))
	imgWinHSV := imgWinBGR.Clone() // gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC3)
	imgWinV := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
	imgWinVNorm := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
	gocv.CvtColor(imgWinBGR, &imgWinHSV, gocv.ColorBGRToHSV)
	gocv.ExtractChannel(imgWinHSV, &imgWinV, 2)
	gocv.Normalize(imgWinV, &imgWinVNorm, 0.0, 255.0, gocv.NormMinMax)

	return imgWinVNorm
}

func MatchResult(img gocv.Mat, imgTemplate gocv.Mat) bool {
	img1 := img
	feature1 := ExtractFeature(img1)
	feature2 := imgTemplate

	result := gocv.NewMatWithSize(2, 2, gocv.MatTypeCV32F)
	mask1 := gocv.NewMatWithSize(32, 80, gocv.MatTypeCV32F)
	gocv.Rectangle(&mask1, image.Rect(0, 0, 80, 32), color.RGBA{255, 255, 255, 0}, -1)

	gocv.MatchTemplate(feature1, feature2, &result, gocv.TmSqdiff, mask1)
	// fmt.Println(result.GetFloatAt(0, 0))
	return result.GetFloatAt(0, 0) < 100.0
}
