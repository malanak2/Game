package types

import (
	"Game/App/Graphics"
	"Game/App/Graphics/Objects"
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
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
	// Init local vars
	keymap = make(map[glfw.Key]glfw.Action)

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
	gCtx.Init()
	// Init shader manager
	Objects.InitShaderManager()

	tris := []Objects.Triangle{
		Objects.NewTriangle(Objects.NewColor(1, 0, 0, 1)),
	}
	app := App{AppState{GraphicalHelper{gCtx, []*Objects.Renderable{}}}, config, time.Now(), 0.0, tris}
	//app.appState.gCtx.AddObjectRenderer(&tris[0].Renderable)
	// Wireframe
	if config.Main.wireframe {
	}
	return &app, nil
}

func (a *App) Run() {
	defer a.appState.gCtx.Destroy()
	for !a.appState.gCtx.Window.ShouldClose() {
		glfw.PollEvents()

		// Do Logic
		a.ProcessInput(a.appState.gCtx.Window)

		oldTime := a.lastFrameTime
		a.lastFrameTime = time.Now()
		a.fps = 1 / a.lastFrameTime.Sub(oldTime).Seconds()

		// graphics
		a.appState.gCtx.Render()
	}
}

var wfmode bool

var keymap map[glfw.Key]glfw.Action

func (a *App) ProcessInput(window *glfw.Window) {
	if window.GetKey(glfw.KeyEscape) == glfw.Press {
		window.SetShouldClose(true)
	}
	if window.GetKey(glfw.KeySpace) == glfw.Press {
		if keymap[glfw.KeySpace] == glfw.Release {
			keymap[glfw.KeySpace] = glfw.Press
			ti := Objects.NewTriangle(Objects.NewColor(0, 1, 0, 1))
			a.tris = append(a.tris, ti)
			a.appState.gCtx.AddObjectRenderer(&ti.Renderable)
		}
	} else {
		keymap[glfw.KeySpace] = glfw.Release
	}
	if window.GetKey(glfw.KeyW) == glfw.Press {
		if keymap[glfw.KeyW] == glfw.Release {
			keymap[glfw.KeyW] = glfw.Press
			if wfmode {
				gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
			} else {
				gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
			}
			wfmode = !wfmode
		}
	} else {
		keymap[glfw.KeyW] = glfw.Release
	}
}
