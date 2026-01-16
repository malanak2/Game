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
}

func (g *GlfwContext) Init() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	g.Window, err = glfw.CreateWindow(1920, 1080, "OpenGL Example", nil, nil)
	if err != nil {
		glfw.Terminate()
		panic(err)
	}

	//g.window.SetKeyCallback(onKey)
	//g.window.SetSizeCallback(onResize)
	g.Window.MakeContextCurrent()

	err = gl.Init()
	if err != nil {
		g.Window.Destroy()
		glfw.Terminate()
		panic(err)
	}

	gl.Viewport(0, 0, 1920, 1080)

	gl.ClearColor(0.0, 0.0, 1.0, 1)
}

func (g *GlfwContext) Destroy() {
	g.Window.Destroy()
	glfw.Terminate()
}
