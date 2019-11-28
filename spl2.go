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

func process_video() {
	win_matched := gocv.NewWindow("Matched")
	win_scoreboard := gocv.NewWindow("Scoreboard")

	img1 := utils.IMRead720p("result.jpg")
	feature1 := scoreboard.Extract_feature(img1)

	var webcam *gocv.VideoCapture
	var err error
	if len(os.Args) == 1 {
		webcam, err = gocv.VideoCaptureDevice(0)
	} else {
		webcam, err = gocv.VideoCaptureFile(os.Args[1])
	}

	fmt.Println(err)
	img := gocv.NewMat()
	img_720p := gocv.NewMatWithSize(1280, 720, gocv.MatTypeCV8UC3)
	for {
		for i := 0; i < 10; i++ {
			webcam.Read(&img)
		}
		// if img.Rows() == 0 {
		//     break
		// }

		gocv.Resize(img, &img_720p, image.Point{1280, 720}, 0.0, 0.0, gocv.InterpolationLinear)
		if scoreboard.Match_result(img_720p, feature1) {
			fmt.Println("Found scoreboard")
			win_scoreboard.IMShow(img_720p)
			win_scoreboard.WaitKey(1)
			return
		}

		if lobby.Match_result(img_720p) {
			fmt.Println("Found lobby_matched")
			win_matched.IMShow(img_720p)
			win_matched.WaitKey(0)
		}
	}
}

func main() {
	img := utils.IMRead720p("masks/spl2_result_scoreboard.jpg")

	if img.Empty() {
		log.Fatal("")
	}

	process_video()
}
