package Objects

type Color struct {
	R, G, B, A float32
}

func NewColor(R float32, G float32, B float32, A float32) Color {
	return Color{R, G, B, A}
}

func (c *Color) ToFloat32Array() []float32 {
	return []float32{c.R, c.G, c.B, c.A}
}
