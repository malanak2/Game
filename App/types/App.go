package types

import (
	"Game/App/Graphics"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type App struct {
	appState AppState
	config   *Config

	lastFrameTime time.Time

	fps float64
}

func InitApp(path *string) (*App, error) {
	config, err := NewConfig(*path)
	if err != nil {
		return nil, err
	}

	//slog.Info("Loading shaders...")
	//
	//prog := loadShaders()
	//
	//slog.Info("Shaders loaded")

	// Create a windowed mode window and its OpenGL context
	//window, err := glfw.CreateWindow(640, 480, "Go Custom Loop", nil, nil)
	//if err != nil {
	//	panic(err)
	//}

	//window.MakeContextCurrent()
	app := App{AppState{GraphicalHelper{&Graphics.GlfwContext{}, []*DrawInfo{}}}, config, time.Now(), 0.0}
	app.appState.gCtx.ctx.Init()

	return &app, nil
}

func (a *App) Run() {
	defer a.appState.gCtx.ctx.Destroy()
	for !a.appState.gCtx.ctx.Window.ShouldClose() {
		glfw.PollEvents()
		// Do Logic
		oldTime := a.lastFrameTime
		a.lastFrameTime = time.Now()
		a.fps = 1 / a.lastFrameTime.Sub(oldTime).Seconds()
		// graphics
		a.appState.gCtx.Render()

		a.appState.gCtx.ctx.Window.SwapBuffers()
	}
}

func loadShaders() uint32 {
	vertexShaders, err := os.ReadDir("./Shaders/Vertex")
	if err != nil {
		panic(err)
	}
	fragmentShaders, err := os.ReadDir("./Shaders/Fragment")
	if err != nil {
		panic(err)
	}
	prog := gl.CreateProgram()
	for i := range vertexShaders {
		slog.Info("Loading shader ./Shaders/Vertex/" + vertexShaders[i].Name())
		fh, err := os.Open("./Shaders/Vertex/" + vertexShaders[i].Name())
		if err != nil {
			slog.Error("Failed to open vertex shader " + vertexShaders[i].Name())
			continue
		}
		file, err := io.ReadAll(fh)
		if err != nil {
			slog.Error("Failed to read all from vertex shader " + vertexShaders[i].Name())
			continue
		}
		shader, err := compileShader(string(file), gl.VERTEX_SHADER)
		if err != nil {
			slog.Error("Failed to compile vertex shader " + vertexShaders[i].Name())
			continue
		}
		gl.AttachShader(prog, shader)
		defer gl.DeleteShader(shader)
	}
	for i := range fragmentShaders {
		slog.Info("Loading shader ./Shaders/Fragment/" + fragmentShaders[i].Name())
		fh, err := os.Open("./Shaders/Fragment/" + fragmentShaders[i].Name())
		if err != nil {
			slog.Error("Failed to open fragment shader " + fragmentShaders[i].Name())
			continue
		}
		file, err := io.ReadAll(fh)
		if err != nil {
			slog.Error("Failed to read all from fragment shader " + fragmentShaders[i].Name())
			continue
		}
		shader, err := compileShader(string(file), gl.FRAGMENT_SHADER)
		if err != nil {
			slog.Error("Failed to compile fragment shader " + fragmentShaders[i].Name())
			continue
		}
		gl.AttachShader(prog, shader)
		defer gl.DeleteShader(shader)
	}
	gl.LinkProgram(prog)
	return prog
}
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	return shader, nil
}
