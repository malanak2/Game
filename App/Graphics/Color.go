package Graphics

type Color struct {
	R, G, B, A float32
}

func (c *Color) ToFloat32Array() []float32 {
	return []float32{c.R, c.G, c.B, c.A}
}
