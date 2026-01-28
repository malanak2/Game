package Graphics

import (
	"github.com/go-gl/gl/v3.3-core/gl"
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

	color Color

	colorLocation int32

	program uint32

	texture *LoadedTexture

	perspLoc int32

	cameraLoc int32

	rotationLoc int32
	Transform   Transform
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
		gl.Uniform4f(r.colorLocation, r.color.R, r.color.G, r.color.B, r.color.A)
	}

	// Apply matrix if got
	if r.perspLoc != -1 {
		// Maybe could optimize - test with more
		gl.UniformMatrix4fv(r.perspLoc, 1, false, &Camera.ProjectionMatrix[0])
		gl.UniformMatrix4fv(r.cameraLoc, 1, false, &Camera.ViewMatrix[0])
	}
	if r.rotationLoc != -1 {
		gl.UniformMatrix4fv(r.rotationLoc, 1, false, &r.Transform.Matrix[0])
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
