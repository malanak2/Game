package Graphics

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type TextRendererT struct {
	TextProgram uint32
	Vao         uint32
	Vbo         uint32
}

var TextRenderer *TextRendererT

func InitTextRenderer() error {
	TextRenderer = &TextRendererT{}
	vert := ShaderManager.LoadVertexShader("text")
	frag := ShaderManager.LoadFragmentShader("text")
	CheckForGLError()
	program := MakeProgram(false, vert, frag)
	gl.UseProgram(program)

	CheckForGLError()
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	CheckForGLError()
	TextRenderer.Vao = vao
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	TextRenderer.Vbo = vbo
	gl.BindVertexArray(TextRenderer.Vao)
	CheckForGLError()

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 6*4*4, nil, gl.DYNAMIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 4, gl.FLOAT, false, 4*4, nil)
	vertI := mgl32.Ortho(0.0, 1920.0, 0.0, 1080, 0, 100)
	gl.UniformMatrix4fv(gl.GetUniformLocation(program, gl.Str("projection\000")), 1, false, &vertI[0])
	TextRenderer.TextProgram = program
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BindVertexArray(0)

	CheckForGLError()

	return nil
}

func (r *TextRendererT) RenderText(text string, x, y float32, scale float32, color Color, fontName string) error {
	gl.UseProgram(TextRenderer.TextProgram)

	gl.Uniform3f(gl.GetUniformLocation(TextRenderer.TextProgram, gl.Str("textColor\x00")), color.R, color.G, color.B)
	textLoc := gl.GetUniformLocation(TextRenderer.TextProgram, gl.Str("text\x00"))
	// 2. Tell the shader: "Look at Texture Unit 0"
	gl.Uniform1i(textLoc, 0)

	// 3. Make sure you are actually using Texture Unit 0
	gl.ActiveTexture(gl.TEXTURE0)
	gl.ActiveTexture(gl.TEXTURE0)

	gl.BindVertexArray(r.Vao)

	// Iterate through all characters
	for _, c := range text {
		ch, ok := FontMgr.Characters[fontName][c]
		if !ok {
			//slog.Warn("Character not found in font", "char", string(c), "font", fontName)
			continue
		}
		var xpos = x + float32(ch.Bearing[0])*scale
		ypos := y - float32(ch.Size[1]-ch.Bearing[1])*scale

		w := float32(ch.Size[0]) * scale
		h := float32(ch.Size[1]) * scale
		// update VBO for each character
		vertices := []float32{
			xpos, ypos + h, 0.0, 0.0,
			xpos, ypos, 0.0, 1.0,
			xpos + w, ypos, 1.0, 1.0,

			xpos, ypos + h, 0.0, 0.0,
			xpos + w, ypos, 1.0, 1.0,
			xpos + w, ypos + h, 1.0, 0.0,
		}

		// render glyph texture over quad
		gl.BindTexture(gl.TEXTURE_2D, ch.TextureID)
		CheckForGLError()
		// update content of VBO memory
		gl.BindBuffer(gl.ARRAY_BUFFER, r.Vbo)
		CheckForGLError()
		gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(vertices)*4, gl.Ptr(&vertices[0])) // be sure to use glBufferSubData and not glBufferData
		CheckForGLError()

		gl.BindBuffer(gl.ARRAY_BUFFER, 0)
		CheckForGLError()
		// render quad
		gl.DrawArrays(gl.TRIANGLES, 0, 6)
		CheckForGLError()
		// now advance cursors for next glyph (note that advance is number of 1/64 pixels)
		x += float32(ch.Advance) * scale // bitshift by 6 to get value in pixels (2^6 = 64 (divide amount of 1/64th pixels by 64 to get amount of pixels))

		CheckForGLError()
	}
	gl.BindVertexArray(0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	return nil
}
