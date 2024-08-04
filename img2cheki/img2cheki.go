package img2cheki

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

type ImageFormat string

const (
	JPEG ImageFormat = "jpeg"
	PNG  ImageFormat = "png"
)

func Img2Cheki(paths []string, output_path string, config Config, output_size Size, format ImageFormat) {
	outputs := make([]*image.RGBA, len(paths))
	for index, path := range paths {
		// 画像を読み込み、チェキのinstax miniサイズの画像に変換
		img := LoadImage(path)
		chekinized := img.ToCheki(config, config.BorderWidth)
		outputs[index] = chekinized
	}

	// 出力画像の配置を計算
	var vertical, horizontal int
	cheki_width := outputs[0].Bounds().Size().X
	cheki_height := outputs[0].Bounds().Size().Y

	// 縦向きにして敷き詰めるときの数
	put1 := (output_size.Width.Pixel() / cheki_width) * (output_size.Height.Pixel() / cheki_height)
	// 横向きにして敷き詰めるときの数
	put2 := (output_size.Width.Pixel() / cheki_height) * (output_size.Height.Pixel() / cheki_width)
	if put1 > put2 {
		vertical = output_size.Height.Pixel()
		horizontal = output_size.Width.Pixel()
	} else {
		vertical = output_size.Width.Pixel()
		horizontal = output_size.Height.Pixel()
	}
	if vertical < cheki_height+config.OutputMargin.Pixel() || horizontal < cheki_width+config.OutputMargin.Pixel() {
		log.Fatal("Output size is too small")
	}

	output_index := 1

	// チェキ画像を一枚ずつ出力画像に配置していく
	for index := 0; index < len(paths); {
		// 出力画像を生成
		output := image.NewRGBA(image.Rect(0, 0, horizontal, vertical))
		FillIn(output, color.White)
		// 横方向に敷き詰めてから、縦方向に敷き詰めていく
		for y := config.OutputMargin.Pixel(); y <= vertical-cheki_height; y += cheki_height + config.OutputMargin.Pixel() {
			for x := config.OutputMargin.Pixel(); x <= horizontal-cheki_width; x += cheki_width + config.OutputMargin.Pixel() {
				if index >= len(paths) {
					break
				}
				// チェキ画像を出力画像に配置
				rect := image.Rect(x, y, output.Rect.Size().X, output.Rect.Size().Y)
				draw.Draw(output, rect, outputs[index], outputs[index].Bounds().Min, draw.Over)
				if index < len(paths) {
					index++
				}
			}
		}
		// jpegファイルとして保存する。ファイル名はoutput_path + index + ".format"
		out, err := os.Create(output_path + fmt.Sprint(output_index) + "." + string(format))
		output_index++
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		switch format {
		case PNG:
			png.Encode(out, output)
		case JPEG:
			jpeg.Encode(out, output, nil)
		}
	}
}
