package scoreboard

import (
//    "fmt"
    "gocv.io/x/gocv"
    "image"
    "image/color"
)

func Extract_feature(img gocv.Mat) gocv.Mat {
    img_win_bgr := img.Region(image.Rect(720, 80, 720 + 80, 80 + 32))
    img_win_hsv := img_win_bgr.Clone() // gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC3)
    img_win_v   := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
    img_win_v_norm := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
    gocv.CvtColor(img_win_bgr, &img_win_hsv, gocv.ColorBGRToHSV)
    gocv.ExtractChannel(img_win_hsv, &img_win_v, 2)
    gocv.Normalize(img_win_v, &img_win_v_norm, 0.0, 255.0, gocv.NormMinMax)

    return img_win_v_norm
}

func Match_result(img gocv.Mat, img_template gocv.Mat) bool {
    img1 := img
    feature1 := Extract_feature(img1)
    feature2 := img_template

    result := gocv.NewMatWithSize(2, 2, gocv.MatTypeCV32F)
    mask_1 := gocv.NewMatWithSize(32, 80, gocv.MatTypeCV32F)
    gocv.Rectangle(&mask_1,image.Rect(0, 0, 80, 32), color.RGBA{255, 255, 255,0}, -1)

    gocv.MatchTemplate(feature1, feature2, &result, gocv.TmSqdiff, mask_1)
    //fmt.Println(result.GetFloatAt(0, 0))
    return result.GetFloatAt(0, 0) < 100.0
}
