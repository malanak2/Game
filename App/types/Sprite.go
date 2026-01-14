package types

import (
	"image"
)

type Sprite struct {
	resourcePath string

	image *image.Image
}
