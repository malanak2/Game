package Graphics

var (
	SQUAREVertices = []float32{
		// Positions 3	// Texture Coord 2 Note: Texture coords scale from 0 to 1 and verticies from -1 to 1 for some reason
		-1, -1, 0, 0, 1, // BL
		1, -1, 0, 1, 1, // BR
		-1, 1, 0, 0, 0, // TL
		1, 1, 0, 1, 0, // TR
	}
	SQUAREIndices = []uint32{
		0, 1, 2,
		1, 2, 3,
	}

	CharacterVertices = []float32{
		// Positions   // Texture Coords
		-0.25, -1, 0, 1,
		0.25, -1, 0, 1,
		-0.25, 1, 0, 1,
		0.25, 1, 0, 1,
	}
	CharacterIndices = []uint32{
		0, 1, 2,
		1, 2, 3,
	}
)

type ObjectManagerT struct {
	Objects []*Triangle
}

var ObjectManager ObjectManagerT

func InitObjectManager() {
	ObjectManager = ObjectManagerT{make([]*Triangle, 0)}
}
