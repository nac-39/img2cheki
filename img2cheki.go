package img2cheki

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"golang.org/x/image/draw"
)

func Img2Cheki(path string, output_path string, config Config, output_size Size) {
	img := LoadImage(path)
	chekinized := img.ToCheki(config, config.FrameThickness)

	// output_size := Size{
	// 	Width:  &Cm{value: 12.7, config: config},
	// 	Height: &Cm{value: 8.9, config: config},
	// }

	output := image.NewRGBA(image.Rect(0, 0, output_size.Width.Pixel(), output_size.Height.Pixel()))
	FillIn(output, color.White)
	draw.Draw(output, output.Bounds(), chekinized, chekinized.Bounds().Min, draw.Over)
	out, err := os.Create(output_path)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	jpeg.Encode(out, output, nil)
}
