package types

import "github.com/go-gl/glfw/v3.3/glfw"

type AppStateT struct {
	isPaused  bool
	lastFrame float64
	DeltaTime float64
	Fps       int
}

var AppState = AppStateT{false, 0, 0, 0}

func (a *AppStateT) updateDeltaTime() {
	currentFrame := glfw.GetTime()
	a.DeltaTime = (currentFrame - a.lastFrame) / 1000
	a.lastFrame = currentFrame
}

func (a *AppStateT) calculateFps() {
	a.Fps = int(1.0 / a.DeltaTime * 1000)
}

func (a *AppStateT) Tick() {
	a.updateDeltaTime()
	a.calculateFps()
}
