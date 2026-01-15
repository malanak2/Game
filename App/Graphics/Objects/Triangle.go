package Objects

import "github.com/go-gl/gl/v4.6-core/gl"

type Triangle struct {
	Renderable
}

func NewTriangle(c Color) Triangle {
	r := Renderable{}
	vertex := LoadVertexShader(`#version 460 core
	layout (location = 0) in vec3 aPos;
	void main()
	{
	   gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
	}`)

	fragment := LoadFragmentShader(`#version 460 core
	out vec4 FragColor;

	void main()
	{
		FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
	} `)

	r.program = MakeProgram(vertex, fragment)

	gl.UseProgram(r.program)

	gl.GenVertexArrays(1, &r.vao)

	gl.BindVertexArray(r.vao)

	r.vertices = []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0}

	_ = LoadVertices(r.vertices)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, nil)
	gl.EnableVertexAttribArray(0)

	return Triangle{}
}
