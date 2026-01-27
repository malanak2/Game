package Graphics

import (
	"errors"
	"fmt"
	"image"
	"image/draw"
	_ "image/png"
	"log/slog"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Texture struct {
	path string
	img  *image.Image
}

type LoadedTexture struct {
	Texture
	id   uint32
	uses int
}

type TextureManager_t struct {
	loadedTextures []LoadedTexture
}

var TextureManager *TextureManager_t

func InitTextureManager() {
	TextureManager = &TextureManager_t{}
	TextureManager.loadedTextures = make([]LoadedTexture, 0)
}

func (l *LoadedTexture) LoadImg() error {
	var er error
	er = nil
	if l.path == "" {
		return errors.New("no path specified")
	}
	if l.img != nil {
		return errors.New("image already loaded")
	}
	reader, err := os.Open(l.path)
	if err != nil {
		reader, err = os.Open("Resources/Textures/Fallback.png")
		slog.Warn("Opening fallback texture", "original", l.path)
		er = errors.New("fallback")
		if err != nil {
			return err
		}
	}
	defer reader.Close()
	img, _, err := image.Decode(reader)
	if err != nil {
		return err
	}
	l.img = &img
	return er
}

func (l *LoadedTexture) UnloadImg() {
	l.uses--
	if l.uses == 0 {
		l.img = nil
		l.id = 0
	}
}

func (l *LoadedTexture) GetImage() (*image.Image, error) {
	if l.img == nil {
		return nil, errors.New("image not loaded")
	}
	l.uses++
	return l.img, nil
}

func (t *TextureManager_t) loadTexture(path string) (*LoadedTexture, error) {
	var texture LoadedTexture
	gl.GenTextures(1, &texture.id)

	gl.BindTexture(gl.TEXTURE_2D, texture.id)

	texture.path = path
	err := texture.LoadImg()
	if err != nil {
		if err.Error() == "fallback" {
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
			gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		} else {
			return nil, err
		}
	}

	rgba := image.NewRGBA((*texture.img).Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, fmt.Errorf("unsupported stride")
	}
	draw.Draw(rgba, rgba.Bounds(), *texture.img, image.Point{0, 0}, draw.Src)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	t.loadedTextures = append(t.loadedTextures, texture)
	return &t.loadedTextures[len(t.loadedTextures)-1], nil
}

func (t *TextureManager_t) GetTexture(path string) (*LoadedTexture, error) {
	for _, texture := range t.loadedTextures {
		if texture.path == path {
			_ = texture.LoadImg()
			return &texture, nil
		}
	}
	texture, err := t.loadTexture(path)
	if err != nil {
		return nil, err
	}
	return texture, nil
}

func (t *TextureManager_t) UnloadTexture(img *image.Image) {
	for _, texture := range t.loadedTextures {
		if texture.img == img {
			texture.UnloadImg()
		}
	}
}
func (t *TextureManager_t) UnloadTextureByPath(imgPath string) {
	for _, texture := range t.loadedTextures {
		if texture.path == imgPath {
			texture.UnloadImg()
		}
	}
}
