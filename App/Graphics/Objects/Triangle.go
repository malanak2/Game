package Objects

import (
	"fmt"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Triangle struct {
	Renderable
}

func NewTriangle(c Color) Triangle {
	r := Renderable{}
	vertex := ShaderManager.LoadVertexShader(`#version 460 core
	layout (location = 0) in vec3 aPos;
	void main()
	{
	   gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
	}`)

	fragment := ShaderManager.LoadFragmentShader(`#version 460 core
	out vec4 FragColor;
	uniform vec4 inCol;	

	void main()
	{
		FragColor = inCol;
	} `)

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
