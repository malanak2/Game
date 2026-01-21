package Graphics

import (
	_ "embed"
	"log/slog"

	"github.com/malanak2/Game/App/config"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

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

	// Debug
	if config.Cfg.Dev.DebugMode {
		glfw.WindowHint(glfw.OpenGLDebugContext, glfw.True)
	}
	g.Window, err = glfw.CreateWindow(1920, 1080, "OpenGL Example", nil, nil)
	if err != nil {
		slog.Error("Failed to create GLFW window", "err", err)
		glfw.Terminate()
		return err
	}
	//g.window.SetKeyCallback(onKey)
	//g.window.SetSizeCallback(onResize):
	g.Window.MakeContextCurrent()
	g.Window.SetFramebufferSizeCallback(FramebufferSizeCallback)

	err = gl.Init()
	if err != nil {
		slog.Error("Failed to init GL", "err", err)
		g.Window.Destroy()
		glfw.Terminate()
		return err
	}
	if !config.Cfg.Main.Vsync {
		glfw.SwapInterval(0)
	} else {
		glfw.SwapInterval(1)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	slog.Info("OpenGL version", "v", version)

	gl.Viewport(0, 0, 1920, 1080)

	gl.ClearColor(0.1, 0.1, 0.1, 1)
	gl.Enable(gl.BLEND)
	gl.Enable(gl.DEPTH_TEST)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

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
	_ = GraphicalManager.Render(-1, false)
	Camera.UpdateScreen(width, height)
}
