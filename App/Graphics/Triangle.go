package Graphics

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Triangle struct {
	Renderable
}

func NewTriangle(c Color) Triangle {
	r := Renderable{}
	vertex := ShaderManager.LoadVertexShader(`basicColor`)

	fragment := ShaderManager.LoadFragmentShader(`basicColor`)

	r.program = MakeProgram(vertex, fragment)

	gl.GenVertexArrays(1, &r.vao)

	gl.BindVertexArray(r.vao)

	cStr := gl.Str("inCol\000")
	r.colorLocation = gl.GetUniformLocation(r.program, cStr)
	fmt.Println(*cStr)

	r.color = c

	r.vertices = []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0}

	r.vbo = ShaderManager.LoadVertices(r.vertices)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	return Triangle{r}
}
func NewTriangleTextured(path string) Triangle {
	r := Renderable{}
	path = "Resources/" + path
	vertex := ShaderManager.LoadVertexShader(`basicTexture`)

	fragment := ShaderManager.LoadFragmentShader(`basicTexture`)

	r.program = MakeProgram(vertex, fragment)

	gl.GenVertexArrays(1, &r.vao)

	gl.BindVertexArray(r.vao)

	cStr := gl.Str("inCol\000")
	r.colorLocation = gl.GetUniformLocation(r.program, cStr)
	fmt.Println(*cStr)
	var err error
	r.texture, err = TextureManager.GetTexture(path)
	if err != nil {
		panic(err)
	}

	r.vertices = SQUAREVertices

	r.indices = SQUAREIndices

	r.vbo, r.ebo = ShaderManager.LoadVerticesWithIndices(r.vertices, r.indices)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	// texture coord attribute
	// Index = location
	// Size = how much
	// Type
	// idk
	// 8 * sizeof(float32)
	// Gde (sum vsech predeslejch size * sizeof(float32)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	return Triangle{r}
}
