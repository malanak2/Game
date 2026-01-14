package types

type DrawInfo struct {
	Vector
	rotation float64

	scale float64

	material Material
}

func (d *DrawInfo) X() int {
	return d.x
}

func (d *DrawInfo) SetX(x int) {
	d.x = x
}

func (d *DrawInfo) Y() int {
	return d.y
}

func (d *DrawInfo) SetY(y int) {
	d.y = y
}

func (d *DrawInfo) Size() float64 {
	return d.scale
}

func (d *DrawInfo) SetSize(size float64) {
	d.scale = size
}

func (d *DrawInfo) Material() Material {
	return d.material
}

func (d *DrawInfo) SetMaterial(material Material) {
	d.material = material
}
