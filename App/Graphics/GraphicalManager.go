package Graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
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

func (ctx *GraphicalmanagerT) Render() error {
	// Clear buffer
	gl.Clear(gl.COLOR_BUFFER_BIT)
	ctx.DrawBackground()

	// Draw objects
	var err error
	for i := range ctx.objects {
		err = (*ctx.objects[i]).Draw()
		if err != nil {
			return err
		}
	}

	ctx.Window.SwapBuffers()
	return nil
}
