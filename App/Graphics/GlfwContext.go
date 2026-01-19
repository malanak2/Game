package Graphics

import (
	_ "embed"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

//go:embed Fonts/Roboto-VariableFont_wdth,wght.ttf
var RobotoVariable []byte

type GlfwContext struct {
	Window *glfw.Window

	Programs []uint32

	Vao []uint32

	width, height int32
}

func (g *GlfwContext) Init() error {
	err := glfw.Init()
	if err != nil {
		return err
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	g.Window, err = glfw.CreateWindow(1920, 1080, "OpenGL Example", nil, nil)
	if err != nil {
		glfw.Terminate()
		return err
	}

	//g.window.SetKeyCallback(onKey)
	//g.window.SetSizeCallback(onResize):
	g.Window.MakeContextCurrent()
	g.Window.SetFramebufferSizeCallback(FramebufferSizeCallback)

	err = gl.Init()
	if err != nil {
		g.Window.Destroy()
		glfw.Terminate()
		return err
	}

	gl.Viewport(0, 0, 1920, 1080)

	gl.ClearColor(0.0, 0.0, 0.0, 1)
	// Disable texture repeating?
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	//gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	return nil
}

func (g *GlfwContext) Destroy() {
	g.Window.Destroy()
	glfw.Terminate()
}

// FramebufferSizeCallback Called when resizing window
func FramebufferSizeCallback(w *glfw.Window, width int, height int) {
	gl.Viewport(0, 0, int32(width), int32(height))
	// Ignore errors, not like we can do much here anyway, better to wait for the proper event loop error bubbling
	_ = GraphicalManager.Render()
}
