package Graphics

import (
	"github.com/malanak2/Game/App/config"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Triangle struct {
	Renderable
}

func NewTriangle(c Color) Triangle {
	r := Renderable{}
	vertex := ShaderManager.LoadVertexShader(`basicColor`)

	fragment := ShaderManager.LoadFragmentShader(`basicColor`)

	r.program = MakeProgram(true, vertex, fragment)

	gl.GenVertexArrays(1, &r.vao)

	gl.BindVertexArray(r.vao)

	cStr := gl.Str("inCol\000")
	r.colorLocation = gl.GetUniformLocation(r.program, cStr)

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
	path = "Resources/Textures/" + path
	vertex := ShaderManager.LoadVertexShader(`basicTexture`)

	fragment := ShaderManager.LoadFragmentShader(`basicTexture`)

	r.program = MakeProgram(true, vertex, fragment)
	gl.UseProgram(r.program)

	r.colorLocation = gl.GetUniformLocation(r.program, gl.Str("inCol\000"))

	r.matrixLoc = gl.GetUniformLocation(r.program, gl.Str("transform\x00"))

	r.perspLoc = gl.GetUniformLocation(r.program, gl.Str("projection\000"))

	r.cameraLoc = gl.GetUniformLocation(r.program, gl.Str("camera\x00"))

	r.rotationLoc = gl.GetUniformLocation(r.program, gl.Str("rotation\x00"))

	r.Rotation = mgl32.NewVecNFromData([]float32{0, 0, 0}).Vec3()

	r.Translation = mgl32.NewVecNFromData([]float32{0, 0, 0}).Vec3()
	MatPerspective := mgl32.Perspective(mgl32.DegToRad(config.Cfg.Main.Fov), float32(1920)/1080, 0.1, 100)
	gl.UniformMatrix4fv(r.perspLoc, 1, false, &MatPerspective[0])
	gl.GenVertexArrays(1, &r.vao)

	gl.BindVertexArray(r.vao)

	var err error
	r.texture, err = TextureManager.GetTexture(path)
	if err != nil {
		panic(err)
	}

	r.matrix = mgl32.Ident4()

	r.scale = 0.5

	r.vertices = SQUAREVertices

	r.indices = SQUAREIndices

	// Binds vbo, ebo
	r.vbo, r.ebo = ShaderManager.LoadVerticesWithIndices(r.vertices, r.indices)

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	// texture coord attribute
	// Index = location
	// Size = how much
	// Type
	// idk
	// 8 * sizeof(float32)
	// Kde (sum vsech predeslejch size * sizeof(float32)
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	CheckForGLError()
	return Triangle{r}
}
