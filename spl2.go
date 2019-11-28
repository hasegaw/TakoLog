package main

import (
	"fmt"
	"log"
	"os"

	"gocv.io/x/gocv"
	"image"

	"github.com/hasegaw/TakoLog/src/scenes/lobby"
	"github.com/hasegaw/TakoLog/src/scenes/result/scoreboard"
	"github.com/hasegaw/TakoLog/src/utils"
)

func processVideo() {
	winMatched := gocv.NewWindow("Matched")
	winScoreboard := gocv.NewWindow("Scoreboard")

	img1 := utils.IMRead720p("result.jpg")
	feature1 := scoreboard.ExtractFeature(img1)

	var webcam *gocv.VideoCapture
	var err error
	if len(os.Args) == 1 {
		webcam, err = gocv.VideoCaptureDevice(0)
	} else {
		webcam, err = gocv.VideoCaptureFile(os.Args[1])
	}

	fmt.Println(err)
	img := gocv.NewMat()
	img720p := gocv.NewMatWithSize(1280, 720, gocv.MatTypeCV8UC3)
	for {
		for i := 0; i < 10; i++ {
			webcam.Read(&img)
		}
		// if img.Rows() == 0 {
		//     break
		// }

		gocv.Resize(img, &img720p, image.Point{1280, 720}, 0.0, 0.0, gocv.InterpolationLinear)
		if scoreboard.MatchResult(img720p, feature1) {
			fmt.Println("Found scoreboard")
			winScoreboard.IMShow(img720p)
			winScoreboard.WaitKey(1)
			return
		}

		if lobby.MatchResult(img720p) {
			fmt.Println("Found lobby_matched")
			winMatched.IMShow(img720p)
			winMatched.WaitKey(0)
		}
	}
}

func main() {
	img := utils.IMRead720p("masks/spl2_result_scoreboard.jpg")

	if img.Empty() {
		log.Fatal("")
	}

	processVideo()
}
