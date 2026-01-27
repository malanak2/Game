package Graphics

import (
	"log/slog"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IRenderable interface {
	Draw() error
}

type Renderable struct {
	vertices []float32

	indices []uint32

	vao uint32

	vbo uint32

	ebo uint32

	x, y, z float32

	scale float32

	color Color

	colorLocation int32

	program uint32

	texture *LoadedTexture

	matrix mgl32.Mat4

	matrixLoc int32

	perspLoc int32

	cameraLoc int32

	Translation mgl32.Vec3

	rotationLoc int32
	// Degrees
	Rotation mgl32.Vec3
}

func (r *Renderable) Draw() error {

	gl.UseProgram(r.program)
	if r.texture != nil {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, r.texture.id)
	}
	gl.BindVertexArray(r.vao)

	// Set color if got
	if r.colorLocation != -1 {
		slog.Info("Setting color")
		gl.Uniform4f(r.colorLocation, r.color.R, r.color.G, r.color.B, r.color.A)
	}

	// Apply matrix if got
	if r.matrixLoc != -1 {
		// Maybe could optimize - test with more
		matrix := r.matrix.Mul4(mgl32.Scale3D(r.scale, r.scale, r.scale))
		gl.UniformMatrix4fv(r.matrixLoc, 1, false, &matrix[0])
		gl.UniformMatrix4fv(r.perspLoc, 1, false, &Camera.ProjectionMatrix[0])
		gl.UniformMatrix4fv(r.cameraLoc, 1, false, &Camera.ViewMatrix[0])
	}
	if r.rotationLoc != -1 {
		mat := mgl32.Translate3D(r.Translation.X(), r.Translation.Y(), r.Translation.Z())
		mat = mat.Mul4(mgl32.Rotate3DX(mgl32.DegToRad(r.Rotation.X())).Mat4())
		mat = mat.Mul4(mgl32.Rotate3DY(mgl32.DegToRad(r.Rotation.Y())).Mat4())
		mat = mat.Mul4(mgl32.Rotate3DZ(mgl32.DegToRad(r.Rotation.Z())).Mat4())
		gl.UniformMatrix4fv(r.rotationLoc, 1, false, &mat[0])
	}

	if r.ebo != 0 {
		gl.DrawElements(gl.TRIANGLES, int32(len(r.indices)), gl.UNSIGNED_INT, nil)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(r.vertices)))
	}

	CheckForGLError()

	gl.BindVertexArray(0)
	return nil
}

func (r *Renderable) Render(shouldRender bool) {
	if shouldRender {
		GraphicalManager.AddObjectRenderer(r)
	} else {
		GraphicalManager.RemoveObjectRenderer(r)
	}
}
