package types

type Vector struct {
	x, y int
}

func (v *Vector) Add(v2 Vector) {
	v.x += v2.x
	v.y += v2.y
}
