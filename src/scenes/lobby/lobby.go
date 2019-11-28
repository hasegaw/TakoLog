package lobby

import (
//    "fmt"
    "gocv.io/x/gocv"
    "image"
    "image/color"

    "utils"
)

var feat_matched gocv.Mat
var done = false


func Extract_feature(img gocv.Mat) gocv.Mat {
    img_win_bgr := img.Region(image.Rect(848, 32, 848+200, 32+32))
    img_win_hsv := img_win_bgr.Clone() // gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC3)
    img_win_v   := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
    img_win_v_norm := gocv.NewMatWithSize(80, 32, gocv.MatTypeCV8UC1)
    gocv.CvtColor(img_win_bgr, &img_win_hsv, gocv.ColorBGRToHSV)
    gocv.ExtractChannel(img_win_hsv, &img_win_v, 2)
    gocv.Normalize(img_win_v, &img_win_v_norm, 0.0, 255.0, gocv.NormMinMax)

    return img_win_v_norm
}

func Match_result(img gocv.Mat) bool {
    if done {
        return false
    }

    img1 := img
    feature1 := Extract_feature(img1)
    feature2 := feat_matched

    result := gocv.NewMatWithSize(1, 1, gocv.MatTypeCV32F)
    mask_1 := gocv.NewMatWithSize(feature1.Rows(), feature1.Cols(), gocv.MatTypeCV32F)
    /*
    fmt.Println(feature1.Cols())
    fmt.Println(feature1.Rows())
    fmt.Println(feature1.Channels())
    fmt.Println(feature2.Cols())
    fmt.Println(feature2.Rows())
    fmt.Println(feature2.Channels())
    fmt.Println(mask_1.Cols())
    fmt.Println(mask_1.Rows())
    fmt.Println(mask_1.Channels())
*/
    gocv.Rectangle(&mask_1, image.Rect(0, 0, mask_1.Cols(), mask_1.Rows()), color.RGBA{255, 255, 255,0}, -1)

    gocv.MatchTemplate(feature1, feature2, &result, gocv.TmSqdiff, mask_1)
    //fmt.Print("lobby_matched: ")
    //fmt.Print(result.GetFloatAt(0, 0))

    done = true

    return result.GetFloatAt(0, 0) < 100.0
}

func init() {
    img := utils.IMRead720p("masks/spl2_lobby_matched.png")
    feat_matched = Extract_feature(img)
}
