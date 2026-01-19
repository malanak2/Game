package Graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
)

type GraphicalManager_t struct {
	*GlfwContext

	objects []*IRenderable
}

var GraphicalManager *GraphicalManager_t

func InitGraphicalManager() error {
	GraphicalManager = &GraphicalManager_t{}
	GraphicalManager.GlfwContext = &GlfwContext{}
	err := GraphicalManager.GlfwContext.Init()
	return err
}

func (ctx *GraphicalManager_t) DrawBackground() {
}

func (ctx *GraphicalManager_t) AddObjectRenderer(object IRenderable) {
	ctx.objects = append(ctx.objects, &object)
}

func (ctx *GraphicalManager_t) RemoveObjectRenderer(object IRenderable) {
	for i := range ctx.objects {
		if ctx.objects[i] == &object {
			ctx.objects = append(ctx.objects[:i], ctx.objects[i+1:]...)
		}
	}
}

func (ctx *GraphicalManager_t) Render() error {
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
