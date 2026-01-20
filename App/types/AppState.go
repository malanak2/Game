package types

import (
	"container/list"

	"github.com/go-gl/glfw/v3.3/glfw"
)

type AppStateT struct {
	isPaused   bool
	lastFrame  float64
	DeltaTime  float64
	Fps        int
	FpsHistory *list.List
	ShowFps    bool
}

var AppState = AppStateT{false, 0, 0, 0, list.New().Init(), false}

func (a *AppStateT) updateDeltaTime() {
	currentFrame := glfw.GetTime()
	a.DeltaTime = (currentFrame - a.lastFrame) / 1000
	a.lastFrame = currentFrame
}

func (a *AppStateT) calculateFps() {
	fps := int(1.0 / (a.DeltaTime * 1000))
	a.FpsHistory.PushBack(fps)
	if a.FpsHistory.Len() > 10 {
		a.FpsHistory.Remove(a.FpsHistory.Front())
	}
	total := 0
	for i := a.FpsHistory.Front(); i != nil; i = i.Next() {
		total += i.Value.(int)
	}
	if a.FpsHistory.Len() == 0 {
		return
	}
	a.Fps = total / a.FpsHistory.Len()
}

func (a *AppStateT) Tick() {
	a.updateDeltaTime()
	a.calculateFps()
}
