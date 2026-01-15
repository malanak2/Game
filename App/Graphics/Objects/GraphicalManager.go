package Objects

import (
	"Game/App/Graphics"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type GraphicalManager_t struct {
	*Graphics.GlfwContext

	objects []*IRenderable
}

var GraphicalManager *GraphicalManager_t

func InitGraphicalManager() {
	GraphicalManager = &GraphicalManager_t{}
	GraphicalManager.GlfwContext = &Graphics.GlfwContext{}
	GraphicalManager.GlfwContext.Init()
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

func (ctx *GraphicalManager_t) Render() {
	// Clear buffer
	gl.Clear(gl.COLOR_BUFFER_BIT)
	ctx.DrawBackground()

	// Draw objects
	for i := range ctx.objects {
		(*ctx.objects[i]).Draw()
	}
	ctx.Window.SwapBuffers()
	glfw.PollEvents()
}
