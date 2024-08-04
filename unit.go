package img2cheki

type Config struct {
	DPI            int
	FrameThickness Unit
	OutputMargin   Unit
}
type Pixel struct {
	value int
	DPI   int
}

func (px *Pixel) Pixel() int {
	return px.value
}

func (px *Pixel) Inch() float64 {
	return float64(px.value) / float64(px.DPI)
}

func (px *Pixel) Cm() float64 {
	return px.Inch() * 2.54
}

func (cm *Cm) Pixel() int {
	return int(float64(cm.Inch()) * float64(cm.DPI))
}

func (cm *Cm) Inch() float64 {
	return cm.value * 2.54
}
func (cm *Cm) Cm() float64 {
	return cm.value
}

type Cm struct {
	value float64
	DPI   int
}

type Unit interface {
	Pixel() int
	Inch() float64
	Cm() float64
}
type Size struct {
	Height, Width Unit
}
