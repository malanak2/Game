package Graphics

import (
	"crypto/sha256"
	"io"
	"log/slog"
	"os"
	"strconv"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type t_ShaderManager struct {
	shaders  map[string]uint32
	programs map[string]uint32
}

var ShaderManager t_ShaderManager

func InitShaderManager() {
	ShaderManager = t_ShaderManager{make(map[string]uint32), make(map[string]uint32)}
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

func (t *t_ShaderManager) LoadVerticesWithIndices(vertices []float32, indices []uint32) (uint32, uint32) {
	vbo := t.LoadVertices(vertices)
	var ebo uint32
	gl.GenBuffers(1, &ebo)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	return vbo, ebo
}

func (t *t_ShaderManager) LoadVertexes(vertices []Vertex, indices []uint32) (uint32, uint32, uint32) {
	//// create buffers/arrays
	var vao uint32
	var vbo uint32
	var ebo uint32
	//glGenVertexArrays(1, &VAO);
	gl.GenVertexArrays(1, &vao)
	//glGenBuffers(1, &VBO);
	gl.GenBuffers(1, &vbo)
	//glGenBuffers(1, &EBO);
	gl.GenBuffers(1, &ebo)
	//
	//glBindVertexArray(VAO);
	gl.BindVertexArray(vao)
	//// load data into vertex buffers
	//glBindBuffer(GL_ARRAY_BUFFER, VBO);
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	//// A great thing about structs is that their memory layout is sequential for all its items.
	//// The effect is that we can simply pass a pointer to the struct and it translates perfectly to a glm::vec3/2 array which
	//// again translates to 3/2 floats which translates to a byte array.
	//glBufferData(GL_ARRAY_BUFFER, vertices.size() * sizeof(Vertex), &vertices[0], GL_STATIC_DRAW);
	slog.Info("Sizeof vertex", "size", unsafe.Sizeof(Vertex{}))
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(Vertex{})), gl.Ptr(vertices), gl.STATIC_DRAW)
	//
	//glBindBuffer(GL_ELEMENT_ARRAY_BUFFER, EBO);
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	//glBufferData(GL_ELEMENT_ARRAY_BUFFER, indices.size() * sizeof(unsigned int), &indices[0], GL_STATIC_DRAW);
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)
	//
	//// set the vertex attribute pointers
	//// vertex Positions
	//glEnableVertexAttribArray(0);
	gl.EnableVertexAttribArray(0)
	//glVertexAttribPointer(0, 3, GL_FLOAT, GL_FALSE, sizeof(Vertex), (void*)0);
	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, int32(unsafe.Sizeof(Vertex{})), 0)
	//// vertex normals
	//glEnableVertexAttribArray(1);
	gl.EnableVertexAttribArray(1)
	//glVertexAttribPointer(1, 3, GL_FLOAT, GL_FALSE, sizeof(Vertex), (void*)offsetof(Vertex, Normal));
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, int32(unsafe.Sizeof(Vertex{})), unsafe.Offsetof(vertices[0].Normal))
	//// vertex texture coords
	//glEnableVertexAttribArray(2);
	gl.EnableVertexAttribArray(2)
	//glVertexAttribPointer(2, 2, GL_FLOAT, GL_FALSE, sizeof(Vertex), (void*)offsetof(Vertex, TexCoords));
	gl.VertexAttribPointerWithOffset(2, 2, gl.FLOAT, false, int32(unsafe.Sizeof(Vertex{})), unsafe.Offsetof(vertices[0].TexCoords))
	//// vertex tangent
	//glEnableVertexAttribArray(3);
	gl.EnableVertexAttribArray(3)
	//glVertexAttribPointer(3, 3, GL_FLOAT, GL_FALSE, sizeof(Vertex), (void*)offsetof(Vertex, Tangent));
	gl.VertexAttribPointerWithOffset(3, 3, gl.FLOAT, false, int32(unsafe.Sizeof(Vertex{})), unsafe.Offsetof(vertices[0].Tangent))
	//// vertex bitangent
	//glEnableVertexAttribArray(4);
	gl.EnableVertexAttribArray(4)
	//glVertexAttribPointer(4, 3, GL_FLOAT, GL_FALSE, sizeof(Vertex), (void*)offsetof(Vertex, Bitangent));
	gl.VertexAttribPointerWithOffset(4, 3, gl.FLOAT, false, int32(unsafe.Sizeof(Vertex{})), unsafe.Offsetof(vertices[0].BitAgent))
	//// ids
	//glEnableVertexAttribArray(5);
	gl.EnableVertexAttribArray(5)
	//glVertexAttribIPointer(5, 4, GL_INT, sizeof(Vertex), (void*)offsetof(Vertex, m_BoneIDs));
	gl.VertexAttribIPointerWithOffset(5, 4, gl.INT, int32(unsafe.Sizeof(Vertex{})), unsafe.Offsetof(vertices[0].MBoneids))
	//
	//// weights
	//glEnableVertexAttribArray(6);
	gl.EnableVertexAttribArray(6)
	//glVertexAttribPointer(6, 4, GL_FLOAT, GL_FALSE, sizeof(Vertex), (void*)offsetof(Vertex, m_Weights));
	gl.VertexAttribPointerWithOffset(6, 4, gl.FLOAT, false, int32(unsafe.Sizeof(Vertex{})), unsafe.Offsetof(vertices[0].MWeights))
	//glBindVertexArray(0);
	gl.BindVertexArray(0)
	return vao, vbo, ebo
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

func (t *t_ShaderManager) checkPCache(shaders ...uint32) (uint32, []byte) {
	first := sha256.New()
	text := ""
	for _, x := range shaders {
		text += strconv.Itoa(int(x))
		text += ","
	}
	first.Write([]byte(text))
	hash := first.Sum(nil)
	if v, ok := t.programs[string(hash)]; ok {
		return v, hash
	}
	return 0, hash
}

func (t *t_ShaderManager) addToPCache(hash []byte, id uint32) {
	t.programs[string(hash)] = id
}

func (t *t_ShaderManager) RemoveFromCache(shaderId uint32) {
	for i, shader := range t.shaders {
		if shader == shaderId {
			delete(t.shaders, i)
		}
	}
}

func (t *t_ShaderManager) LoadVertexShader(shaderSource string) uint32 {
	shaderSource = "Resources/Shaders/" + shaderSource + ".vertex"
	v, hash := t.checkCache(shaderSource)
	if v != 0 {
		slog.Info("Hit cache")
		return v
	}
	// Load File
	file, err := os.Open(shaderSource)
	if err != nil {
		slog.Error("Vertex shader at this path doesnt exist", "path", shaderSource)
		return 0
	}
	defer file.Close()
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	in, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Failed to read from vertex shader", "path", shaderSource)
	}
	sources, freeSources := gl.Strs(string(in) + "\000")
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

func (t *t_ShaderManager) MakeProgram(reusableShaders bool, shaders ...uint32) uint32 {
	v, hash := t.checkPCache(shaders...)
	if v != 0 {
		slog.Info("Hit shader cache")
		return v
	}
	program := gl.CreateProgram()
	CheckForGLError()
	for _, shader := range shaders {
		gl.AttachShader(program, shader)
		slog.Info("Attaching shader", "shader", shader)
		if CheckForGLError() {
			slog.Error("Failed to attach shader", "program", program, "shader", shader)
		}
	}

	gl.LinkProgram(program)
	CheckForGLError()
	var (
		success int32
		infoLog [512]uint8
	)
	CheckForGLError()
	gl.GetProgramiv(program, gl.LINK_STATUS, &success)
	if success != gl.TRUE {
		gl.GetProgramInfoLog(program, 512, nil, &infoLog[0])
		panic(gl.GoStr(&infoLog[0]))
	}
	CheckForGLError()

	if !reusableShaders {
		for _, shader := range shaders {
			gl.DeleteShader(shader)
			ShaderManager.RemoveFromCache(shader)
		}
	}
	CheckForGLError()
	t.addToPCache(hash, program)
	return program
}

func (t *t_ShaderManager) LoadFragmentShader(shaderSource string) uint32 {
	shaderSource = "Resources/Shaders/" + shaderSource + ".fragment"
	v, hash := t.checkCache(shaderSource)
	if v != 0 {
		return v
	}
	// Load File
	file, err := os.Open(shaderSource)
	if err != nil {
		slog.Error("Fragment shader at this path doesnt exist", "path", shaderSource)
		return 0
	}
	defer file.Close()
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	in, err := io.ReadAll(file)
	if err != nil {
		slog.Error("Failed to read from fragment shader", "path", shaderSource)
	}
	sources, freeSources := gl.Strs(string(in) + "\000")
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
