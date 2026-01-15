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
	}` + "\000")

	fragment := ShaderManager.LoadFragmentShader(`#version 460 core
	out vec4 FragColor;
	uniform vec4 inCol;	

	void main()
	{
		FragColor = inCol;
	}` + "\000")

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
	vertex := ShaderManager.LoadVertexShader(`#version 460 core
layout (location = 0) in vec3 aPos;
layout (location = 1) in vec2 aTexCoord;
out vec2 TexCoord;
void main()
{
    gl_Position = vec4(aPos, 1.0);
    TexCoord = aTexCoord;
}` + "\000")

	fragment := ShaderManager.LoadFragmentShader(`#version 460 core
	out vec4 FragColor;
	uniform sampler2D textureIn;
	in vec2 TexCoord;
	void main()
	{
		FragColor = texture(textureIn, TexCoord);
	}` + "\000")

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
