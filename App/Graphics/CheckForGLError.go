//go:build dev

package Graphics

import (
	"log/slog"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func CheckForGLError() {
	for code := gl.GetError(); code != gl.NO_ERROR; code = gl.GetError() {
		var err string
		switch code {
		case gl.INVALID_ENUM:
			err = "INVALID_ENUM"
			break
		case gl.INVALID_VALUE:
			err = "INVALID_VALUE"
			break
		case gl.INVALID_OPERATION:
			err = "INVALID_OPERATION"
			break
		case gl.STACK_OVERFLOW:
			err = "STACK_OVERFLOW"
			break
		case gl.STACK_UNDERFLOW:
			err = "STACK_UNDERFLOW"
			break
		case gl.OUT_OF_MEMORY:
			err = "OUT_OF_MEMORY"
			break
		case gl.INVALID_FRAMEBUFFER_OPERATION:
			err = "INVALID_FRAMEBUFFER_OPERATION"
			break
		}
		_, file, line, ok := runtime.Caller(1)
		if ok {
			slog.Error("OpenGL Error", "err", err, "gl", code, "file", file, "line", line)
		} else {
			slog.Error("OpenGL Error, failed to get caller", "err", err, "gl", code)
		}
	}
}
