package types

import (
	"log/slog"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/malanak2/Game/App/Graphics"
	config2 "github.com/malanak2/Game/App/config"

	"github.com/go-gl/gl/v3.3-core/gl"
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

func spawnTexturedCube() error {
	vec := mgl32.NewVecN(3)
	ti := Graphics.NewCube(vec.Vec3(), "trump.png")
	ti.Render(true)
	//a.tris = append(a.tris, ti)
	return nil
}

func toggleFps() error {
	AppState.ShowFps = !AppState.ShowFps
	return nil
}

func toggleVsync() error {
	if config2.Cfg.Main.Vsync {
		glfw.SwapInterval(0)
		config2.Cfg.Main.Vsync = false
	} else {
		glfw.SwapInterval(1)
		config2.Cfg.Main.Vsync = true
	}
	return nil
}

func InitApp(path *string) error {
	err := config2.InitConfig(*path)
	if err != nil {
		return err
	}

	// Initialize managers
	slog.Info("Initializing Graphics Manager")
	err = Graphics.InitGraphicalManager()
	if err != nil {
		slog.Error("Error initializing Graphics Manager")
		return err
	}
	slog.Info("Initializing Shader Manager")
	Graphics.InitShaderManager()
	slog.Info("Initializing Texture Manager")
	Graphics.InitTextureManager()
	slog.Info("Initializing Object Manager")
	Graphics.InitObjectManager()
	slog.Info("Initializing Camera")
	Graphics.InitCamera([3]float32{0, 0, 0}, [3]float32{0, 1, 0})

	slog.Info("Initializing Font Manager")
	Graphics.InitFontManager()
	slog.Info("Initializing Text Renderer")
	err = Graphics.InitTextRenderer()
	if err != nil {
		return err
	}
	InitKeybindManager()

	// Dev Keybinds
	if config2.Cfg.Dev.Dev {
		slog.Info("Initializing Development Mode")
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
		KeybindManager.AddOnPressed(glfw.KeyF3, spawnTexturedCube)
		KeybindManager.AddOnPressed(glfw.KeyF4, toggleFps)
		KeybindManager.AddOnPressed(glfw.KeyF5, toggleVsync)
	}
	Graphics.GraphicalManager.Window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	Graphics.GraphicalManager.Window.SetCursorPosCallback(Graphics.MouseCallback)
	Graphics.GraphicalManager.Window.SetScrollCallback(Graphics.ScrollWheelCallback)
	// Load default font
	slog.Info("Initialization Complete")
	slog.Info("Loading default font")
	err = Graphics.LoadFont("Default")
	if err != nil {
		slog.Error("Error loading default font")
		return err
	}
	wfState = false

	return nil
}

func Run() error {
	defer Graphics.GraphicalManager.Destroy()
	var err error
	slog.Info("Entering main loop")
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
