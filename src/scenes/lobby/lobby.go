package lobby

import (
	// "fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"

	"github.com/hasegaw/TakoLog/src/utils"
)

var featMatched gocv.Mat
var done = false

func ExtractFeature(img gocv.Mat) gocv.Mat {
	imgWinBGR := img.Region(image.Rect(848, 32, 848+200, 32+32))
	imgWinHSV := imgWinBGR.Clone() // gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC3)
	imgWinV := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
	imgWinVNorm := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
    gocv.CvtColor(imgWinBGR, &imgWinHSV, gocv.ColorBGRToHSV)
	gocv.ExtractChannel(imgWinHSV, &imgWinV, 2)
	gocv.Normalize(imgWinV, &imgWinVNorm, 0.0, 255.0, gocv.NormMinMax)

	return imgWinVNorm
}

func MatchResult(img gocv.Mat) bool {
	if done {
		return false
	}

	img1 := img
	feature1 := ExtractFeature(img1)
	feature2 := featMatched

	result := gocv.NewMatWithSize(1, 1, gocv.MatTypeCV32F)
	mask1 := gocv.NewMatWithSize(feature1.Rows(), feature1.Cols(), gocv.MatTypeCV32F)
	// fmt.Println(feature1.Cols())
	// fmt.Println(feature1.Rows())
	// fmt.Println(feature1.Channels())
	// fmt.Println(feature2.Cols())
	// fmt.Println(feature2.Rows())
	// fmt.Println(feature2.Channels())
	// fmt.Println(mask1.Cols())
	// fmt.Println(mask1.Rows())
	// fmt.Println(mask1.Channels())
	gocv.Rectangle(&mask1, image.Rect(0, 0, mask1.Cols(), mask1.Rows()), color.RGBA{255, 255, 255, 0}, -1)

	gocv.MatchTemplate(feature1, feature2, &result, gocv.TmSqdiff, mask1)
	// fmt.Print("lobby_matched: ")
	// fmt.Print(result.GetFloatAt(0, 0))

	done = true

	return result.GetFloatAt(0, 0) < 100.0
}

func init() {
	img := utils.IMRead720p("masks/spl2_lobby_matched.png")
	featMatched = ExtractFeature(img)
}
