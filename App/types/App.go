package types

import (
	"Game/App/Graphics"
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
	Graphics.GraphicalManager.Window.SetShouldClose(true)
}

var wfState bool

func ToggleWireFrame() {
	if wfState {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
	wfState = !wfState
}

func spawnTexturedTriangle() {
	ti := Graphics.NewTriangleTextured("Surprise.png")
	ti.Render(true)
	//a.tris = append(a.tris, ti)
}

func InitApp(path *string) (*App, error) {
	config, err := NewConfig(*path)
	if err != nil {
		return nil, err
	}

	// Initialize managers
	err = Graphics.InitGraphicalManager()
	if err != nil {
		return nil, err
	}
	Graphics.InitShaderManager()
	Graphics.InitTextureManager()
	Graphics.InitObjectManager()
	InitKeybindManager()

	wfState = false

	KeybindManager.AddOnPressed(glfw.KeyEscape, closeApp)
	KeybindManager.AddOnPressed(glfw.KeyW, ToggleWireFrame)
	KeybindManager.AddOnPressed(glfw.KeySpace, spawnTexturedTriangle)

	app := App{config, time.Now(), 0.0}
	return &app, nil
}

func (a *App) Run() error {
	defer Graphics.GraphicalManager.Destroy()
	var err error
	for !Graphics.GraphicalManager.Window.ShouldClose() {
		glfw.PollEvents()

		// Do Logic
		err = KeybindManager.HandleInput(Graphics.GraphicalManager.Window)
		if err != nil {
			return err
		}

		oldTime := a.lastFrameTime
		a.lastFrameTime = time.Now()
		a.fps = 1 / a.lastFrameTime.Sub(oldTime).Seconds()

		// Graphics
		err = Graphics.GraphicalManager.Render()
		if err != nil {
			return err
		}
	}
	return nil
}
