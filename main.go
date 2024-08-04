package main

import "github.com/nac-39/img2cheki/img2cheki"

func main() {
	const dpi = 350
	unit := img2cheki.UnitConfig{DPI: dpi}
	images := []string{"sample1.jpeg", "sample2.jpeg", "sample3.jpeg"}        // example
	output_prefix := "output"                                                 // example
	output_size := img2cheki.Size{Width: unit.Cm(12.7), Height: unit.Cm(8.9)} // 日本のL版サイズ

	img2cheki.Img2Cheki(images, output_prefix,
		img2cheki.Config{DPI: dpi, BorderWidth: unit.Cm(0.01), OutputMargin: unit.Cm(0.1)},
		output_size, img2cheki.JPEG)
}
