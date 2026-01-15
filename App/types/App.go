package types

import (
	"Game/App/Graphics"
	"Game/App/Graphics/Objects"
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type App struct {
	appState AppState
	config   *Config

	lastFrameTime time.Time

	fps float64

	tris []Objects.Triangle
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
	gCtx := &Graphics.GlfwContext{}
	tris := []Objects.Triangle{
		Objects.NewTriangle(Objects.NewColor(1, 0, 0, 1)),
	}
	app := App{AppState{GraphicalHelper{gCtx, []*Drawable{}}}, config, time.Now(), 0.0, tris}
	app.appState.gCtx.ctx.Init()

	return &app, nil
}

func (a *App) Run() {
	defer a.appState.gCtx.ctx.Destroy()
	for !a.appState.gCtx.ctx.Window.ShouldClose() {
		glfw.PollEvents()
		// Do Logic
		a.appState.gCtx.ctx.ProcessInput()
		oldTime := a.lastFrameTime
		a.lastFrameTime = time.Now()
		a.fps = 1 / a.lastFrameTime.Sub(oldTime).Seconds()
		// graphics
		a.appState.gCtx.Render()

		a.appState.gCtx.ctx.Window.SwapBuffers()
	}
}
