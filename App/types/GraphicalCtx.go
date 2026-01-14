package types

import (
	"Game/App/Graphics"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Drawable interface {
	GetDrawInfo() *DrawInfo
}
type GraphicalHelper struct {
	ctx *Graphics.GlfwContext

	objects []*DrawInfo
}

func (ctx *GraphicalHelper) DrawBackground() {
	gl.ClearColor(0.0, 255.0, 0.0, 1.0)
}

func (ctx *GraphicalHelper) AddObjectRender(object Drawable) {
	ctx.objects = append(ctx.objects, object.GetDrawInfo())
}

func (ctx *GraphicalHelper) RemoveObjectRender(object Drawable) {
	info := object.GetDrawInfo()
	for i := range ctx.objects {
		if ctx.objects[i] == info {
			ctx.objects = append(ctx.objects[:i], ctx.objects[i+1:]...)
		}
	}
}

func (ctx *GraphicalHelper) drawObject(info *DrawInfo) {
	if info.Material().image.image != nil {
		//ctx.ctx.DrawRectangle(float64(info.X()), float64(info.Y()), 1*info.scale, 1*info.scale)
	}
	//ctx.ctx.DrawImage(*info.Material().image.image, info.X(), info.Y())
}

func (ctx *GraphicalHelper) Render() {
	//ctx.ctx.Clear()
	gl.ClearColor(0.0, 255.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	//vertices := []float64{
	//	-0.5, -0.5, 0.0,
	//	0.5, -0.5, 0.0,
	//	0.0, 0.5, 0.0,
	//}
	//gl.UseProgram(ctx.program)
	ctx.DrawBackground()
	ctx.ctx.RenderText("Ahoj kamaradi bububu")

	for i := range ctx.objects {
		ctx.drawObject(ctx.objects[i])
	}
}
