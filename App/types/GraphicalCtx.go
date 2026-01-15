package types

import (
	"Game/App/Graphics"
	"Game/App/Graphics/Objects"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Drawable interface {
	Draw()
}
type GraphicalHelper struct {
	*Graphics.GlfwContext

	objects []*Objects.Renderable
}

func (ctx *GraphicalHelper) DrawBackground() {
}

func (ctx *GraphicalHelper) AddObjectRenderer(object *Objects.Renderable) {
	ctx.objects = append(ctx.objects, object)
}

func (ctx *GraphicalHelper) RemoveObjectRenderer(object *Objects.Renderable) {
	for i := range ctx.objects {
		if ctx.objects[i] == object {
			ctx.objects = append(ctx.objects[:i], ctx.objects[i+1:]...)
		}
	}
}

func (ctx *GraphicalHelper) Render() {
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
