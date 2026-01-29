package Graphics

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/malanak2/assimp"
)

func LoadModel(dir, file string) error {
	ex, err := os.Executable()
	if err != nil {
		return err
	}
	exPath := filepath.Dir(ex)
	scene := assimp.ImportFile(exPath+"/Resources/Models/"+dir+"/"+file, uint(assimp.Process_Triangulate /*|assimp.Process_FlipUVs*/))
	if scene == nil || scene.Flags()&assimp.SceneFlags_Incomplete != 0 || scene.RootNode() == nil {
		return errors.New("Assimp error: " + assimp.GetErrorString())
	}
	processNode(scene.RootNode(), scene, dir)
	return nil
}

func processNode(node *assimp.Node, scene *assimp.Scene, dir string) {
	//// process each mesh located at the current node
	//for(unsigned int i = 0; i < node->mNumMeshes; i++)
	for i := range node.NumMeshes() {
		//// the node object only contains indices to index the actual Objects in the scene.
		//// the scene contains all the data, node is just to keep stuff organized (like relations between nodes).
		//aiMesh* mesh = scene->mMeshes[node->mMeshes[i]];
		//meshes.push_back(processMesh(mesh, scene));
		mesh := scene.Meshes()[node.Meshes()[i]]
		meshProcessed := processMesh(mesh, scene, dir)
		meshProcessed.Render(true)
		ObjectManager.PushObject(&meshProcessed)
	}
	//// after we've processed all of the meshes (if any) we then recursively process each of the children nodes
	//for(unsigned int i = 0; i < node->mNumChildren; i++)
	for i := range node.NumChildren() {
		//processNode(node->mChildren[i], scene);
		processNode(node.Children()[i], scene, dir)
	}
}

func processMesh(mesh *assimp.Mesh, scene *assimp.Scene, dir string) Renderable {
	//	// data to fill
	//	vector<Vertex> vertices;
	var vertices []Vertex
	//	vector<unsigned int> indices;
	var indeces []uint32
	//	vector<Texture> textures;
	var textures []*LoadedTexture
	//
	//	// walk through each of the mesh's vertices
	//	for(unsigned int i = 0; i < mesh->mNumVertices; i++)
	for i, vert := range mesh.Vertices() {
		//	Vertex vertex;
		var ver Vertex
		//	glm::vec3 vector; // we declare a placeholder vector since assimp uses its own vector class that doesn't directly convert to glm's vec3 class so we transfer the data to this placeholder glm::vec3 first.
		var vec mgl32.Vec3
		//	// positions
		//	vector.x = mesh->mVertices[i].x;
		//	vector.y = mesh->mVertices[i].y;
		//	vector.z = mesh->mVertices[i].z;
		vec = mgl32.NewVecNFromData([]float32{vert.X(), vert.Y(), vert.Z()}).Vec3()
		//	vertex.Position = vector;
		ver.Position = vec
		//	if (mesh->HasNormals())
		if mesh.Normals() != nil {
			//  vector.x = mesh->mNormals[i].x;
			//	vector.y = mesh->mNormals[i].y;
			//	vector.z = mesh->mNormals[i].z;
			vec = mgl32.NewVecNFromData([]float32{mesh.Normals()[i].X(), mesh.Normals()[i].Y(), mesh.Normals()[i].Z()}).Vec3()
			//	vertex.Normal = vector;
			ver.Normal = vec
		}
		//	// texture coordinates
		//	if(mesh->mTextureCoords[0]) // does the mesh contain texture coordinates?
		if mesh.TextureCoords(0) != nil {
			//	glm::vec2 vec;
			//	// a vertex can contain up to 8 different texture coordinates. We thus make the assumption that we won't
			//	// use models where a vertex can have multiple texture coordinates so we always take the first set (0).
			//	vec.x = mesh->mTextureCoords[0][i].x;
			//	vec.y = mesh->mTextureCoords[0][i].y;
			//	vertex.TexCoords = vec;
			ver.TexCoords = mgl32.NewVecNFromData([]float32{mesh.TextureCoords(0)[i].X(), mesh.TextureCoords(0)[i].Y()}).Vec2()
			//	// tangent
			//	vector.x = mesh->mTangents[i].x;
			//	vector.y = mesh->mTangents[i].y;
			//	vector.z = mesh->mTangents[i].z;
			//	vertex.Tangent = vector;
			if mesh.Tangents() != nil {
				ver.Tangent = mgl32.NewVecNFromData([]float32{mesh.Tangents()[i].X(), mesh.Tangents()[i].Y(), mesh.Tangents()[i].Z()}).Vec3()
			}
			//	// bitangent
			//	vector.x = mesh->mBitangents[i].x;
			//	vector.y = mesh->mBitangents[i].y;
			//	vector.z = mesh->mBitangents[i].z;
			//	vertex.Bitangent = vector;
			if mesh.Bitangents() != nil {
				ver.BitAgent = mgl32.NewVecNFromData([]float32{mesh.Bitangents()[i].X(), mesh.Bitangents()[i].Y(), mesh.Bitangents()[i].Z()}).Vec3()
			}
		} else {
			//	else
			ver.TexCoords = mgl32.NewVecNFromData([]float32{0, 0}).Vec2()
		}
		//	vertices.push_back(vertex);
		vertices = append(vertices, ver)
	}

	//	// now wak through each of the mesh's faces (a face is a mesh its triangle) and retrieve the corresponding vertex indices.
	//	for(unsigned int i = 0; i < mesh->mNumFaces; i++)
	for _, face := range mesh.Faces() {

		//	aiFace face = mesh->mFaces[i];
		//	// retrieve all indices of the face and store them in the indices vector
		//	for(unsigned int j = 0; j < face.mNumIndices; j++)
		for _, ind := range face.CopyIndices() {
			//	indices.push_back(face.mIndices[j]);
			indeces = append(indeces, ind)
		}
	}
	//	// process materials
	//	aiMaterial* material = scene->mMaterials[mesh->mMaterialIndex];
	material := scene.Materials()[mesh.MaterialIndex()]
	//	// we assume a convention for sampler names in the shaders. Each diffuse texture should be named
	//	// as 'texture_diffuseN' where N is a sequential number ranging from 1 to MAX_SAMPLER_NUMBER.
	//	// Same applies to other texture as the following list summarizes:
	//	// diffuse: texture_diffuseN
	//	// specular: texture_specularN
	//	// normal: texture_normalN
	//
	//	// 1. diffuse maps
	//	vector<Texture> diffuseMaps = loadMaterialTextures(material, aiTextureType_DIFFUSE, "texture_diffuse");
	diffuseMaps := loadMaterialTextures(material, assimp.TextureType_Diffuse, "texture_diffuse", dir)
	//	textures.insert(textures.end(), diffuseMaps.begin(), diffuseMaps.end());
	textures = append(textures, diffuseMaps...)
	//	// 2. specular maps
	//	vector<Texture> specularMaps = loadMaterialTextures(material, aiTextureType_SPECULAR, "texture_specular");
	specularMaps := loadMaterialTextures(material, assimp.TextureType_Specular, "texture_specular", dir)
	//	textures.insert(textures.end(), specularMaps.begin(), specularMaps.end());,
	textures = append(textures, specularMaps...)
	//	// 3. normal maps
	//std::vector<Texture> normalMaps = loadMaterialTextures(material, aiTextureType_HEIGHT, "texture_normal");
	normalMaps := loadMaterialTextures(material, assimp.TextureType_Height, "texture_normal", dir)
	//	textures.insert(textures.end(), normalMaps.begin(), normalMaps.end());
	textures = append(textures, normalMaps...)
	//	// 4. height maps
	//std::vector<Texture> heightMaps = loadMaterialTextures(material, aiTextureType_AMBIENT, "texture_height");
	heightMaps := loadMaterialTextures(material, assimp.TextureType_Ambient, "texture_height", dir)
	//	textures.insert(textures.end(), heightMaps.begin(), heightMaps.end());
	textures = append(textures, heightMaps...)
	//
	//	// return a mesh object created from the extracted mesh data
	//	return Mesh(vertices, indices, textures);
	rend := Renderable{vertexes: vertices, indices: indeces, textures: textures}
	vertexshader := ShaderManager.LoadVertexShader("model")
	fragmentshader := ShaderManager.LoadFragmentShader("model")
	vao, vbo, ebo := ShaderManager.LoadVertexes(vertices, indeces)
	program := ShaderManager.MakeProgram(true, vertexshader, fragmentshader)
	gl.UseProgram(program)
	rend.program = program
	rend.colorLocation = gl.GetUniformLocation(rend.program, gl.Str("inCol\000"))

	rend.perspLoc = gl.GetUniformLocation(rend.program, gl.Str("projection\000"))

	rend.cameraLoc = gl.GetUniformLocation(rend.program, gl.Str("camera\x00"))

	rend.rotationLoc = gl.GetUniformLocation(rend.program, gl.Str("rotation\x00"))
	rend.Transform.Scale = 1
	rend.Transform.SetPos(mgl32.NewVecNFromData([]float32{0, 0, 0}).Vec3())
	rend.Transform.Rotation = mgl32.NewVecNFromData([]float32{0, 0, 0}).Vec3()
	rend.Transform.updateMatrix()
	rend.vao = vao
	rend.vbo = vbo
	rend.ebo = ebo
	return rend
}

// vector<Texture> loadMaterialTextures(aiMaterial *mat, aiTextureType type, string typeName)
func loadMaterialTextures(mat *assimp.Material, kind assimp.TextureType, kindName, dir string) []*LoadedTexture {
	//vector<Texture> textures;
	textures := make([]*LoadedTexture, 0)
	//for(unsigned int i = 0; i < mat->GetTextureCount(type); i++)
	for i := range mat.GetMaterialTextureCount(kind) {
		//aiString str;
		//mat->GetTexture(type, i, &str);
		str, _, _, _, _, _, _, _ := mat.GetMaterialTexture(kind, i)
		// This is already implemented in TextureManager, so we use that
		//// check if texture was loaded before and if so, continue to next iteration: skip loading a new texture
		//bool skip = false;
		//for(unsigned int j = 0; j < textures_loaded.size(); j++)
		//{
		//if(std::strcmp(textures_loaded[j].path.data(), str.C_Str()) == 0)
		//{
		//textures.push_back(textures_loaded[j]);
		//skip = true; // a texture with the same filepath has already been loaded, continue to next one. (optimization)
		//break;
		//}
		//}
		//if(!skip)
		//{   // if texture hasn't been loaded already, load it
		//Texture texture;
		//texture.id = TextureFromFile(str.C_Str(), this->directory);
		//texture.type = typeName;
		//texture.path = str.C_Str();
		//textures.push_back(texture);
		//textures_loaded.push_back(texture);  // store it as texture loaded for entire model, to ensure we won't unnecessary load duplicate textures.
		//}
		//}
		tex, err := TextureManager.GetTexture(dir + "/" + str)
		if err != nil {
			slog.Error("Failed to load material texture", "path", str)
			return nil
		}
		tex.name = kindName
		textures = append(textures, tex)
	}
	return textures
}
