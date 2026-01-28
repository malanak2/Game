package Graphics

import (
	"strconv"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/malanak2/Game/App/config"
)

type GraphicalmanagerT struct {
	*GlfwContext

	objects []*IRenderable
}

var GraphicalManager *GraphicalmanagerT

func InitGraphicalManager() error {
	GraphicalManager = &GraphicalmanagerT{}
	GraphicalManager.GlfwContext = &GlfwContext{}
	err := GraphicalManager.GlfwContext.Init()
	return err
}

func (ctx *GraphicalmanagerT) DrawBackground() {
}

func (ctx *GraphicalmanagerT) AddObjectRenderer(object IRenderable) {
	ctx.objects = append(ctx.objects, &object)
}

func (ctx *GraphicalmanagerT) RemoveObjectRenderer(object IRenderable) {
	for i := range ctx.objects {
		if ctx.objects[i] == &object {
			ctx.objects = append(ctx.objects[:i], ctx.objects[i+1:]...)
		}
	}
}

func (ctx *GraphicalmanagerT) Render(fps int, showFps bool) error {
	// Clear buffer
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.Clear(gl.DEPTH_BUFFER_BIT)
	ctx.DrawBackground()

	// Draw objects
	var err error
	for i := range ctx.objects {
		err = (*ctx.objects[i]).Draw()
		CheckForGLError()
		if err != nil {
			return err
		}
	}
	if showFps {
		// Draw FPS
		err = TextRenderer.RenderText(strconv.Itoa(fps)+" fps", 100, 540, 0.5, Color{1, 1, 1, 1}, "Default")
		CheckForGLError()
		if err != nil {
			return err
		}
	}

	if config.Cfg.Dev.Dev {
		err = TextRenderer.RenderText("F1: Toggle wireframe, F2: Spawn textured triangle, F3: Spawn textured cube, F4: Toggle FPS, F5: Toggle VSync, Escape: Quit", 10, 1000, 0.33, Color{1, 1, 1, 1}, "Default")
	}

	ctx.Window.SwapBuffers()
	return nil
}
