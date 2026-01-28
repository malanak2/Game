package Graphics

import "github.com/go-gl/mathgl/mgl32"

type Transform struct {
	Translation, Rotation mgl32.Vec3
	Scale                 float32

	Matrix mgl32.Mat4
}

func NewTransform(Pos, Rot mgl32.Vec3) Transform {
	return NewScaledTransform(Pos, Rot, 1)
}
func NewScaledTransform(Pos, Rot mgl32.Vec3, Scale float32) Transform {
	t := Transform{Translation: Pos, Rotation: Rot, Scale: Scale}
	t.updateMatrix()
	return t
}

func (t *Transform) SetPos(Position mgl32.Vec3) {
	t.Translation = Position
	t.updateMatrix()
}

func (t *Transform) MovePos(Position mgl32.Vec3) {
	t.Translation = t.Translation.Add(Position)
	t.updateMatrix()
}

func (t *Transform) RotateX(Angle float32) {
	t.Rotation = t.Rotation.Add(mgl32.NewVecNFromData([]float32{Angle, 0, 0}).Vec3())
	t.updateMatrix()
}

func (t *Transform) RotateY(Angle float32) {
	t.Rotation = t.Rotation.Add(mgl32.NewVecNFromData([]float32{0, Angle, 0}).Vec3())
	t.updateMatrix()
}

func (t *Transform) RotateZ(Angle float32) {
	t.Rotation = t.Rotation.Add(mgl32.NewVecNFromData([]float32{0, 0, Angle}).Vec3())
	t.updateMatrix()
}

func (t *Transform) updateMatrix() {
	mat := mgl32.Scale3D(t.Scale, t.Scale, t.Scale)
	mat = mat.Mul4(mgl32.Translate3D(t.Translation.X(), t.Translation.Y(), t.Translation.Z()))
	mat = mat.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(t.Rotation.X())).Mat4())
	mat = mat.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(t.Rotation.Y())).Mat4())
	mat = mat.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(t.Rotation.Z())).Mat4())
	t.Matrix = mat
}
