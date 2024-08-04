package img2cheki

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

type ImgOperator interface {
	ToCheki(config Config, frame_thickness Unit) *image.RGBA
	Save()
}

type GoImg struct {
	// image
	Image image.Image

	// file path
	Path string

	// height and width of image
	Height, Width int
}

func LoadImage(path string) GoImg {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	src, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	size := src.Bounds().Size()
	width, height := size.X, size.Y

	img := GoImg{
		Image:  src,
		Path:   path,
		Height: height,
		Width:  width,
	}

	return img
}

func (img *GoImg) Save(path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Println("Cannot create file:", err)
	}
	defer file.Close()

	png.Encode(file, img.Image)
}

func ResizeKeepAspect(img image.Image, width, height int) *image.RGBA {
	src := img.Bounds()
	dst := image.Rect(0, 0, width, height)
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))
	left := (width - src.Dx()) / 2
	top := (height - src.Dy()) / 2
	draw.BiLinear.Scale(newImage, newImage.Bounds(), img, image.Rect(left, top, width, height), draw.Over, nil)
	resized := image.NewRGBA(dst)
	draw.CatmullRom.Scale(resized, dst, img, src, draw.Over, nil)
	return resized
}

func FillIn(img *image.RGBA, color color.Color) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, color)
		}
	}
}

func AddBlackFrame(img *image.RGBA, thickness int) *image.RGBA {
	newImage := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx()+thickness*2, img.Bounds().Dy()+thickness*2))
	FillIn(newImage, color.Black)
	draw.Draw(newImage, image.Rect(thickness, thickness, newImage.Rect.Size().X, newImage.Rect.Size().Y), img, img.Bounds().Min, draw.Over)
	return newImage
}

func (img *GoImg) ToCheki(config Config, frame_thickness Unit) *image.RGBA {

	resized_size := Size{
		Width:  &Cm{value: 4.6, config: config},
		Height: &Cm{value: 6.2, config: config},
	}

	resized := ResizeKeepAspect(img.Image, resized_size.Width.Pixel(), resized_size.Height.Pixel())

	frame_size := Size{
		Width:  &Cm{value: 5.4, config: config},
		Height: &Cm{value: 8.6, config: config},
	}
	frame := image.NewRGBA(image.Rect(0, 0, frame_size.Width.Pixel(), frame_size.Height.Pixel()))
	FillIn(frame, color.White)
	frame = AddBlackFrame(frame, frame_thickness.Pixel())

	margin := (frame_size.Width.Pixel() + frame_thickness.Pixel()*2 - resized_size.Width.Pixel()) / 2
	draw.Draw(frame, image.Rect(margin, margin, frame.Rect.Size().X, frame.Rect.Size().Y), resized, resized.Bounds().Min, draw.Over)
	return frame
}
