package types

import (
	"Game/App/Graphics"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Drawable interface {
	Draw()
}
type GraphicalHelper struct {
	ctx *Graphics.GlfwContext

	objects []*Drawable
}

func (ctx *GraphicalHelper) DrawBackground() {
}

func (ctx *GraphicalHelper) AddObjectRenderer(object *Drawable) {
	ctx.objects = append(ctx.objects, object)
}

func (ctx *GraphicalHelper) RemoveObjectRenderer(object *Drawable) {
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
}
