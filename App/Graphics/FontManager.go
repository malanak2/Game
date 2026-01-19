package Graphics

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Character struct {
	TextureID uint32
	Size      [2]int32
	Bearing   [2]int32
	Advance   uint32
}

type CharactersT map[rune]Character

type FontManager struct {
	Characters map[string]CharactersT

	textProgram uint32
}

var FontMgr FontManager

func InitFontManager() {
	FontMgr = FontManager{make(map[string]CharactersT), 0}
}

func LoadFont(fontPath string) error {
	og := fontPath
	fontPath = "Resources/Fonts/" + fontPath + ".ttf"
	_, exists := FontMgr.Characters[og]
	if exists {
		return nil
	}
	FontMgr.Characters[og] = CharactersT{}
	reader, err := os.Open(fontPath)
	if err != nil {
		return fmt.Errorf("failed to open font file: %v", err)
	}
	defer reader.Close()
	file, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("failed to read font file: %v", err)
	}
	f, err := freetype.ParseFont(file)
	if err != nil {
		return fmt.Errorf("failed to load font: %v", err)
	}

	face := truetype.NewFace(f, &truetype.Options{Size: 48})
	defer face.Close()
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	for c := rune(0); c < 128; c++ {
		// Would probably be better to skip null term hmhmmm
		if c == '\000' {
			continue
		}
		bounds, advance, ok := face.GlyphBounds(c)
		if !ok {
			log.Printf("Failed to get glyph bounds for character: %c", c)
			continue
		}
		width := (bounds.Max.X - bounds.Min.X).Ceil()
		height := (bounds.Max.Y - bounds.Min.Y).Ceil()
		if width == 0 || height == 0 {
			slog.Info("Skipping character with no visible representation", "char", string(c))
			continue
		}
		rgba := image.NewAlpha(image.Rect(0, 0, width, height))
		dot := fixed.Point26_6{
			X: -bounds.Min.X,
			Y: -bounds.Min.Y,
		}

		drawer := font.Drawer{
			Dst:  rgba,
			Src:  image.NewUniform(color.Alpha{A: 255}),
			Face: face,
			Dot:  dot,
		}
		drawer.DrawString(string(c))
		var texture uint32
		gl.GenTextures(1, &texture)
		gl.BindTexture(gl.TEXTURE_2D, texture)

		rgbaV := drawer.Dst.(*image.Alpha)
		gl.TexImage2D(
			gl.TEXTURE_2D,
			0,
			gl.RED,
			int32(width),
			int32(height),
			0,
			gl.RED,
			gl.UNSIGNED_BYTE,
			gl.Ptr(rgbaV.Pix),
		)

		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

		character := Character{
			TextureID: texture,
			Size:      [2]int32{int32(width), int32(height)},
			Bearing:   [2]int32{int32(bounds.Min.X.Floor()), int32(-bounds.Min.Y.Floor())},
			Advance:   uint32(advance.Floor()),
		}

		CheckForGLError()

		slog.Info("Loaded character", "char", string(c), "size", character.Size, "bearing", character.Bearing, "advance", character.Advance, "texID", character.TextureID, "font", fontPath)

		FontMgr.Characters[og][c] = character
	}
	slog.Info("Loaded font", "font", fontPath, "test", FontMgr.Characters[fontPath])

	gl.BindTexture(gl.TEXTURE_2D, 0)
	return nil
}
