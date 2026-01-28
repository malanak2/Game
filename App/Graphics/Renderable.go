package Graphics

import (
	"strconv"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IRenderable interface {
	Draw() error
}

type Vertex struct {
	Position  mgl32.Vec3
	Normal    mgl32.Vec3
	TexCoords mgl32.Vec2
	Tangent   mgl32.Vec3
	BitAgent  mgl32.Vec3
	MBoneids  [4]int
	//weights from each bone
	MWeights [4]float32
}

type Renderable struct {
	vertices []float32

	indices []uint32

	vao uint32

	vbo uint32

	ebo uint32

	color Color

	colorLocation int32

	program uint32

	textures []*LoadedTexture

	perspLoc int32

	cameraLoc int32

	rotationLoc int32
	Transform
	vertexes []Vertex
}

type RenderableOptions struct {
	UseVertices bool
	UseIndices  bool
	IsTextured  bool
	IsColored   bool
}

func NewRenderable(options RenderableOptions) Renderable {
	panic("Not implemented")
	return Renderable{}
}

func (r *Renderable) Draw() error {
	diffuseNr := 1
	specularNr := 1

	gl.UseProgram(r.program)
	if r.textures != nil {
		for i, tex := range r.textures {
			gl.ActiveTexture(uint32(gl.TEXTURE0 + i))
			gl.BindTexture(gl.TEXTURE_2D, tex.id)
			var number string
			name := tex.name

			if name == "texture_diffuse" {
				number = strconv.Itoa(diffuseNr)
				diffuseNr++
			} else if name == "texture_specular" {
				number = strconv.Itoa(specularNr)
				specularNr++
			}
			gl.Uniform1i(gl.GetUniformLocation(r.program, gl.Str("material."+name+number+"\x00")), int32(i))
			gl.BindTexture(gl.TEXTURE_2D, tex.id)
		}
		gl.ActiveTexture(gl.TEXTURE0)
		CheckForGLError()
	}
	gl.BindVertexArray(r.vao)

	CheckForGLError()
	// Set color if got
	if r.colorLocation != -1 {
		gl.Uniform4f(r.colorLocation, r.color.R, r.color.G, r.color.B, r.color.A)
	}

	CheckForGLError()
	// Apply matrix if got
	if r.perspLoc != -1 {
		// Maybe could optimize - test with more
		gl.UniformMatrix4fv(r.perspLoc, 1, false, &Camera.ProjectionMatrix[0])
		gl.UniformMatrix4fv(r.cameraLoc, 1, false, &Camera.ViewMatrix[0])
	}
	if r.rotationLoc != -1 {
		gl.UniformMatrix4fv(r.rotationLoc, 1, false, &r.Transform.Matrix[0])
	}

	CheckForGLError()
	if r.ebo != 0 {
		gl.DrawElements(gl.TRIANGLES, int32(len(r.indices)), gl.UNSIGNED_INT, nil)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(r.vertices)))
	}

	if CheckForGLError() {
		//slog.Error("Failed to render object", "vertexes", r.vertexes)
	}

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
