package Graphics

import (
	"Game/App/config"
	"log/slog"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type IRenderable interface {
	Draw() error
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

	matrix mgl32.Mat4

	matrixLoc int32

	perspLocation int32

	cameraLocation int32
}

var (
	MatPerspective = mgl32.Perspective(mgl32.DegToRad(config.Cfg.Main.Fov), float32(1920)/1080, 0.1, 100)
	camera         = mgl32.LookAtV(mgl32.Vec3{3, 3, 3}, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0})
)

func (r *Renderable) Draw() error {
	gl.UseProgram(r.program)
	if r.texture != nil {
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, r.texture.id)
	}
	gl.BindVertexArray(r.vao)

	// Set color if got
	if r.colorLocation != -1 {
		slog.Info("Setting color")
		gl.Uniform4f(r.colorLocation, r.color.R, r.color.G, r.color.B, r.color.A)
	}

	// Apply matrix if got
	if r.matrixLoc != -1 {
		//r.matrix = r.matrix.Mul4(mgl32.Scale3D(r.scale, r.scale, 1))
		gl.UniformMatrix4fv(r.matrixLoc, 1, false, &r.matrix[0])
	}

	if r.indices != nil {
		gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, nil)
	} else {
		gl.DrawArrays(gl.TRIANGLES, 0, int32(len(r.vertices)))
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
