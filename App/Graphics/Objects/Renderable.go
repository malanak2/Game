package Objects

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type IRenderable interface {
	Draw()
}

type Renderable struct {
	vertices []float32

	indices []int32

	vao uint32

	vbo uint32

	ebo uint32

	x, y, z float32

	scale float32

	color Color

	colorLocation int32

	program uint32

	texture *LoadedTexture
}

func (r *Renderable) Draw() {
	gl.UseProgram(r.program)
	if r.texture != nil {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, r.texture.id)
	}
	gl.BindVertexArray(r.vao)

	if r.texture != nil {
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	} else {
		gl.Uniform4f(r.colorLocation, r.color.R, r.color.G, r.color.B, r.color.A)

		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(r.vertices)))
	}

	gl.BindVertexArray(0)
}

func (r *Renderable) Render(shouldRender bool) {
	if shouldRender {
		GraphicalManager.AddObjectRenderer(r)
	} else {
		GraphicalManager.RemoveObjectRenderer(r)
	}
}
