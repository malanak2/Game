package Graphics

import (
	"errors"
	"fmt"
	"image"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Character struct {
	Size    [2]int32
	Bearing [2]int32
	Advance uint32
	TR, BL  mgl32.Vec2
}

type CharactersT map[rune]Character

type FontManager struct {
	Characters map[string]CharactersT

	textProgram map[string]uint32
}

var FontMgr FontManager

func InitFontManager() {
	FontMgr = FontManager{make(map[string]CharactersT), make(map[string]uint32)}
}

func LoadFont(fontPath string) error {
	og := fontPath
	// modify the path so that it is Where we want it + this way we can omit the ttf since we dont support other formats anyway
	fontPath = "Resources/Fonts/" + fontPath + ".ttf"
	// Check if we cached the font yet
	_, exists := FontMgr.Characters[og]
	if exists {
		return nil
	}

	// Initialize the array
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

	face := truetype.NewFace(f, &truetype.Options{Size: 64})
	defer face.Close()
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	// Atlas width, height
	w := 0
	h := 0

	for c := rune(32); c < 128; c++ {
		if c == '\000' {
			continue
		}
		bounds, _, _ := face.GlyphBounds(c)
		width := (bounds.Max.X - bounds.Min.X).Ceil()
		height := (bounds.Max.Y - bounds.Min.Y).Ceil()
		if c == ' ' {
			width = 20
			height = 26
		}
		w += width + 2
		if h < height {
			h = height
		}
	}

	var atlasId uint32
	gl.GenTextures(1, &atlasId)
	gl.BindTexture(gl.TEXTURE_2D, atlasId)

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RED,
		int32(w),
		int32(h),
		0,
		gl.RED,
		gl.UNSIGNED_BYTE,
		gl.Ptr(nil),
	)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	x := 0

	// Characters in the ascii table < 32 are control characters that arent renderable anyway, no need to load them
	for c := rune(32); c < 128; c++ {
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
		if (width == 0 || height == 0) && c != ' ' {
			slog.Info("Skipping character with no visible representation", "char", string(c))
			continue
		}
		if c == ' ' {
			width = 20
			height = 26
		}
		rgba := image.NewAlpha(image.Rect(0, 0, width, height))
		dot := fixed.Point26_6{
			X: -bounds.Min.X,
			Y: -bounds.Min.Y,
		}

		drawer := font.Drawer{
			Dst:  rgba,
			Src:  image.Opaque,
			Face: face,
			Dot:  dot,
		}
		drawer.DrawString(string(c))
		rgbaV := drawer.Dst.(*image.Alpha)
		gl.TextureSubImage2D(
			atlasId,
			0,
			int32(x),
			0,
			int32(width),
			int32(height),
			gl.RED,
			gl.UNSIGNED_BYTE,
			gl.Ptr(rgbaV.Pix),
		)
		if CheckForGLError() {
			slog.Error("Failed", "x", x, "width", width, "height", height, "atlas width", w, "atlas height", h, "char", string(c))
			return errors.New("Failed")
		}

		bearingX := int32(bounds.Min.X.Floor())
		bearingY := int32(-bounds.Min.Y.Floor())
		if c == ' ' {
			bearingX = 2
			bearingY = 35
		}
		character := Character{
			Size:    [2]int32{int32(width), int32(height)},
			Bearing: [2]int32{bearingX, bearingY},
			Advance: uint32(advance.Floor()),
			// X, Y
			TR: mgl32.NewVecNFromData([]float32{float32(x+width) / float32(w), float32(height) / float32(h)}).Vec2(),
			// X, Y
			BL: mgl32.NewVecNFromData([]float32{float32(x) / float32(w), 0}).Vec2(),
		}

		CheckForGLError()

		FontMgr.Characters[og][c] = character
		x += width + 2
	}
	slog.Info("Loaded font", "font", fontPath)

	gl.BindTexture(gl.TEXTURE_2D, 0)
	FontMgr.textProgram[og] = atlasId
	return nil
}
