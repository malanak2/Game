package Graphics

import (
	_ "embed"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/pekim/freetype"
)

//go:embed Fonts/Roboto-VariableFont_wdth,wght.ttf
var RobotoVariable []byte

type GlfwContext struct {
	Window *glfw.Window
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

	g.Window, err = glfw.CreateWindow(300, 300, "OpenGL Example", nil, nil)
	if err != nil {
		glfw.Terminate()
		panic(err)
	}

	lib, err := freetype.Init()
	if err != nil {
		panic(err)
	}

	face, err := lib.NewMemoryFace(RobotoVariable, 0)
	if err != nil {
		panic(err)
	}

	err = face.SetPixelSizes(0, 32)
	if err != nil {
		panic(err)
	}

	//g.window.SetKeyCallback(onKey)
	//g.window.SetSizeCallback(onResize)
	g.Window.MakeContextCurrent()
	err = gl.Init()
	if err != nil {
		glfw.Terminate()
		g.Window.Destroy()
		panic(err)
	}
}

func (g *GlfwContext) Destroy() {
	glfw.Terminate()
	g.Window.Destroy()
}
