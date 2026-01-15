package Objects

import "github.com/go-gl/gl/v4.6-core/gl"

func LoadVertices(vertices []float32) uint32 {
	// Generate buffer
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// Float32 has 4 bytes so *4
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(&vertices[0]), gl.STATIC_DRAW)
	return vbo
}

func LoadVerticesWithIndices(vertices []float32, indices []int32) (uint32, uint32) {
	vbo := LoadVertices(vertices)
	var ebo uint32
	// 3. copy our index array in a element buffer for OpenGL to use
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	return vbo, ebo
}

func LoadVertexShader(shaderSource string) uint32 {
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
	return shader
}

func MakeProgram(shaders ...uint32) uint32 {
	program := gl.CreateProgram()
	for shader := range shaders {
		gl.AttachShader(program, shaders[shader])
		defer gl.DeleteShader(shaders[shader])
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

func LoadFragmentShader(shaderSource string) uint32 {
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
	return shader
}
