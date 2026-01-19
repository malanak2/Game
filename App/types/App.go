package types

import (
	"Game/App/Graphics"
	config2 "Game/App/config"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type App struct {
}

func closeApp() error {
	Graphics.GraphicalManager.Window.SetShouldClose(true)
	return nil
}

var wfState bool

func ToggleWireFrame() error {
	if wfState {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
	} else {
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	}
	wfState = !wfState
	return nil
}

func spawnTexturedTriangle() error {
	ti := Graphics.NewTriangleTextured("Surprise.png")
	ti.Render(true)
	//a.tris = append(a.tris, ti)
	return nil
}

func toggleFps() error {
	AppState.ShowFps = !AppState.ShowFps
	return nil
}

func InitApp(path *string) error {
	err := config2.InitConfig(*path)
	if err != nil {
		return err
	}

	// Initialize managers
	err = Graphics.InitGraphicalManager()
	if err != nil {
		return err
	}
	Graphics.InitShaderManager()
	Graphics.InitTextureManager()
	Graphics.InitObjectManager()
	Graphics.InitCamera([3]float32{0, 0, 0}, [3]float32{0, 1, 0})

	Graphics.InitFontManager()
	err = Graphics.InitTextRenderer()
	if err != nil {
		return err
	}
	InitKeybindManager()

	if config2.Cfg.Dev.Dev {
		KeybindManager.AddOnHeld(glfw.KeyW, func() error {
			Graphics.Camera.MoveCamera(Graphics.CameraForward, float32(AppState.DeltaTime))
			return nil
		})
		KeybindManager.AddOnHeld(glfw.KeyS, func() error {
			Graphics.Camera.MoveCamera(Graphics.CameraBackward, float32(AppState.DeltaTime))
			return nil
		})
		KeybindManager.AddOnHeld(glfw.KeyA, func() error {
			Graphics.Camera.MoveCamera(Graphics.CameraLeft, float32(AppState.DeltaTime))
			return nil
		})
		KeybindManager.AddOnHeld(glfw.KeyD, func() error {
			Graphics.Camera.MoveCamera(Graphics.CameraRight, float32(AppState.DeltaTime))
			return nil
		})
		KeybindManager.AddOnHeld(glfw.KeySpace, func() error {
			Graphics.Camera.MoveCamera(Graphics.CameraUp, float32(AppState.DeltaTime))
			return nil
		})
		KeybindManager.AddOnHeld(glfw.KeyLeftShift, func() error {
			Graphics.Camera.MoveCamera(Graphics.CameraDown, float32(AppState.DeltaTime))
			return nil
		})
		KeybindManager.AddOnPressed(glfw.KeyEscape, closeApp)
		KeybindManager.AddOnPressed(glfw.KeyF1, ToggleWireFrame)
		KeybindManager.AddOnPressed(glfw.KeyF2, spawnTexturedTriangle)
		KeybindManager.AddOnPressed(glfw.KeyF3, toggleFps)
	}
	err = Graphics.LoadFont("Default")
	if err != nil {
		return err
	}
	wfState = false

	return nil
}

func Run() error {
	defer Graphics.GraphicalManager.Destroy()
	var err error
	for !Graphics.GraphicalManager.Window.ShouldClose() {
		glfw.PollEvents()

		// Handle Keybinds
		err = KeybindManager.HandleInput(Graphics.GraphicalManager.Window)
		if err != nil {
			return err
		}

		// Calculate Deltatime
		AppState.Tick()
		// Update Camera
		Graphics.Camera.Calculate()

		// Graphics
		err = Graphics.GraphicalManager.Render(AppState.Fps, AppState.ShowFps)
		if err != nil {
			return err
		}
		Graphics.CheckForGLError()
	}
	return nil
}
