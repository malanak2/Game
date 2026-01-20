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

	perspLocation int32

	cameraLocation int32
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
		matrix := r.matrix.Mul4(mgl32.Scale3D(r.scale, r.scale, 1))
		//matrix = mgl32.HomogRotate3DY(float32(glfw.GetTime()))
		gl.UniformMatrix4fv(r.matrixLoc, 1, false, &matrix[0])
		gl.UniformMatrix4fv(r.perspLocation, 1, false, &Camera.ProjectionMatrix[0])
		gl.UniformMatrix4fv(r.cameraLocation, 1, false, &Camera.ViewMatrix[0])
	}

	if r.indices != nil {
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
