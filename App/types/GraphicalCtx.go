package types

import (
	"github.com/fogleman/gg"
)

type Drawable interface {
	GetDrawInfo() *DrawInfo
}
type GraphicalCtx struct {
	ctx *gg.Context

	objects []*DrawInfo
}

func (ctx *GraphicalCtx) DrawBackground() {
	ctx.ctx.SetRGB(255, 255, 255)
}

func (ctx *GraphicalCtx) AddObjectRender(object Drawable) {
	ctx.objects = append(ctx.objects, object.GetDrawInfo())
}

func (ctx *GraphicalCtx) RemoveObjectRender(object Drawable) {
	info := object.GetDrawInfo()
	for i := range ctx.objects {
		if ctx.objects[i] == info {
			ctx.objects = append(ctx.objects[:i], ctx.objects[i+1:]...)
		}
	}
}

func (ctx *GraphicalCtx) drawObject(info *DrawInfo) {
	if info.Material().image.image != nil {
		ctx.ctx.DrawRectangle(float64(info.X()), float64(info.Y()), 1*info.scale, 1*info.scale)
	}
	ctx.ctx.DrawImage(*info.Material().image.image, info.X(), info.Y())
}

func (ctx *GraphicalCtx) Render() {
	ctx.ctx.Clear()
	ctx.DrawBackground()

	for i := range ctx.objects {
		ctx.drawObject(ctx.objects[i])
	}
}
