package types

import (
	"Game/App/Graphics/Objects"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type App struct {
	config *Config

	lastFrameTime time.Time

	fps float64
}

func closeApp() {
	Objects.GraphicalManager.Window.SetShouldClose(true)
}

var wfState bool

func toggleWireFrame() {
	if wfState {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
	wfState = !wfState
}

func spawnTexturedTriangle() {
	ti := Objects.NewTriangleTextured("Surprise.png")
	ti.Render(true)
	//a.tris = append(a.tris, ti)
}

func InitApp(path *string) (*App, error) {
	config, err := NewConfig(*path)
	if err != nil {
		return nil, err
	}
	// InitTextureManager local vars

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

	// InitTextureManager shader manager
	Objects.InitGraphicalManager()
	Objects.InitShaderManager()
	Objects.InitTextureManager()
	Objects.InitObjectManager()
	InitKeybindManager()
	wfState = false
	KeybindManager.AddOnPressed(glfw.KeyEscape, closeApp)
	KeybindManager.AddOnPressed(glfw.KeyW, toggleWireFrame)
	KeybindManager.AddOnPressed(glfw.KeySpace, spawnTexturedTriangle)

	app := App{config, time.Now(), 0.0}
	//app.appState.gCtx.AddObjectRenderer(&tris[0].Renderable)
	// Wireframe
	if config.Main.wireframe {
	}
	return &app, nil
}

func (a *App) Run() {
	defer Objects.GraphicalManager.Destroy()
	for !Objects.GraphicalManager.Window.ShouldClose() {
		glfw.PollEvents()

		// Do Logic
		KeybindManager.HandleInput(Objects.GraphicalManager.Window)

		oldTime := a.lastFrameTime
		a.lastFrameTime = time.Now()
		a.fps = 1 / a.lastFrameTime.Sub(oldTime).Seconds()

		// graphics
		Objects.GraphicalManager.Render()
	}
}
