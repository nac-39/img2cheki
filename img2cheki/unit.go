package img2cheki

type Config struct {
	DPI          int
	BorderWidth  Unit
	OutputMargin Unit
}

type Pixel struct {
	Value int
	DPI   int
}

func (px *Pixel) Pixel() int {
	return px.Value
}

func (px *Pixel) Inch() float64 {
	return float64(px.Value) / float64(px.DPI)
}

func (px *Pixel) Cm() float64 {
	return px.Inch() * 2.54
}

func (cm *Cm) Pixel() int {
	return int(float64(cm.Inch()) * float64(cm.DPI))
}

func (cm *Cm) Inch() float64 {
	return cm.Value * 2.54
}
func (cm *Cm) Cm() float64 {
	return cm.Value
}

type Cm struct {
	Value float64
	DPI   int
}

type Unit interface {
	Pixel() int
	Inch() float64
	Cm() float64
}

type UnitConfig struct {
	DPI int
}

func (uc *UnitConfig) Pixel(value int) *Pixel {
	return &Pixel{Value: value, DPI: uc.DPI}
}

func (uc *UnitConfig) Cm(value float64) *Cm {
	return &Cm{Value: value, DPI: uc.DPI}
}

type Size struct {
	Height, Width Unit
}
