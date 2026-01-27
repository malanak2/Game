package Graphics

import (
	"log/slog"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

func NewCube(coords mgl32.Vec3, texturePath string) *Renderable {
	r := Renderable{}
	texturePath = "Resources/Textures/" + texturePath
	vertex := ShaderManager.LoadVertexShader(`basicTexture`)

	fragment := ShaderManager.LoadFragmentShader(`basicTexture`)

	r.program = MakeProgram(true, vertex, fragment)
	gl.UseProgram(r.program)

	r.colorLocation = gl.GetUniformLocation(r.program, gl.Str("inCol\000"))

	r.matrixLoc = gl.GetUniformLocation(r.program, gl.Str("transform\x00"))

	r.perspLoc = gl.GetUniformLocation(r.program, gl.Str("projection\000"))

	r.cameraLoc = gl.GetUniformLocation(r.program, gl.Str("camera\x00"))

	r.translationLoc = gl.GetUniformLocation(r.program, gl.Str("translation\x00"))

	r.rotationLoc = gl.GetUniformLocation(r.program, gl.Str("rotation\x00"))

	slog.Info("Locs", "translation", r.translationLoc, "rotation", r.rotationLoc)
	gl.UniformMatrix4fv(r.perspLoc, 1, false, &Camera.ProjectionMatrix[0])

	gl.GenVertexArrays(1, &r.vao)

	gl.BindVertexArray(r.vao)

	var err error
	r.texture, err = TextureManager.GetTexture(texturePath)
	if err != nil {
		panic(err)
	}

	r.matrix = mgl32.Ident4()

	r.scale = 0.5

	r.vertices = CUBEVertices

	r.Translation = coords

	r.Rotation = mgl32.NewVecNFromData([]float32{1, 1, 1}).Vec3()

	// Binds vbo, ebo
	r.vbo = ShaderManager.LoadVertices(CUBEVertices)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 5*4, 0)
	gl.EnableVertexAttribArray(0)
	// texture coord attribute
	// Index = location
	// Size = how much
	// Type
	// idk
	// 8 * sizeof(float32)
	// Kde (sum vsech predeslejch size * sizeof(float32)
	gl.VertexAttribPointerWithOffset(1, 2, gl.FLOAT, false, 5*4, 3*4)
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	gl.BindVertexArray(0)

	CheckForGLError()

	return &r
}
