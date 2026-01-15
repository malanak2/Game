package Objects

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type Renderable struct {
	vertices []float32

	indices []int

	vao uint32

	vbo uint32

	program uint32
}

func (r *Renderable) Draw() {
	gl.UseProgram(r.program)
	gl.BindVertexArray(r.vao)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(r.vertices)))
	gl.BindVertexArray(0)
}
