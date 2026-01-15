package Objects

import (
	"crypto/sha256"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type t_ShaderManager struct {
	shaders map[string]uint32
}

var ShaderManager t_ShaderManager

var (
	SHADER_VERTEX_BASIC = ""
)

var (
	SHADER_FRAGMENT_BASIC = ""
)

func InitShaderManager() {
	ShaderManager = t_ShaderManager{make(map[string]uint32)}
}

func (t *t_ShaderManager) LoadVertices(vertices []float32) uint32 {
	// Generate buffer
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// Float32 has 4 bytes so *4
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(&vertices[0]), gl.STATIC_DRAW)
	return vbo
}

func (t *t_ShaderManager) LoadVerticesWithIndices(vertices []float32, indices []int32) (uint32, uint32) {
	vbo := t.LoadVertices(vertices)
	var ebo uint32
	gl.GenBuffers(1, &ebo)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	return vbo, ebo
}

func (t *t_ShaderManager) checkCache(shaderSource string) (uint32, []byte) {
	first := sha256.New()
	first.Write([]byte(shaderSource))

	hash := first.Sum(nil)
	if v, ok := t.shaders[string(hash)]; ok {
		return v, hash
	}
	return 0, hash
}

func (t *t_ShaderManager) addToCache(shaderSource []byte, id uint32) {
	t.shaders[string(shaderSource)] = id
}

func (t *t_ShaderManager) LoadVertexShader(shaderSource string) uint32 {
	v, hash := t.checkCache(shaderSource)
	if v != 0 {
		return v
	}
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	sources, freeSources := gl.Strs(shaderSource)
	defer freeSources()
	gl.ShaderSource(shader, 1, sources, nil)
	gl.CompileShader(shader)
	var (
		success int32
		infoLog [512]uint8
	)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success != gl.TRUE {
		gl.GetShaderInfoLog(shader, 512, nil, &infoLog[0])
		panic(gl.GoStr(&infoLog[0]))
	}
	t.addToCache(hash, shader)
	return shader
}

func MakeProgram(shaders ...uint32) uint32 {
	program := gl.CreateProgram()
	for shader := range shaders {
		gl.AttachShader(program, shaders[shader])
	}
	gl.LinkProgram(program)
	var (
		success int32
		infoLog [512]uint8
	)
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)
	if success != gl.TRUE {
		gl.GetProgramInfoLog(program, 512, nil, &infoLog[0])
		panic(gl.GoStr(&infoLog[0]))
	}
	return program
}

func (t *t_ShaderManager) LoadFragmentShader(shaderSource string) uint32 {

	v, hash := t.checkCache(shaderSource)
	if v != 0 {
		return v
	}

	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	sources, freeSources := gl.Strs(shaderSource)
	defer freeSources()
	gl.ShaderSource(shader, 1, sources, nil)
	gl.CompileShader(shader)
	var (
		success int32
		infoLog [512]uint8
	)
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &success)
	if success != gl.TRUE {
		gl.GetShaderInfoLog(shader, 512, nil, &infoLog[0])
		panic(gl.GoStr(&infoLog[0]))
	}

	t.addToCache(hash, shader)

	return shader
}
