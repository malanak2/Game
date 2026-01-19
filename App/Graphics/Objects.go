package Graphics

var (
	SQUAREVertices = []float32{
		// Positions 3	// Texture Coord 2
		-1.0, -1.0, 0.0, -1.0, 1.0, // BL
		1.0, -1.0, 0.0, 1.0, 1.0, // BR
		-1.0, 1.0, 0.0, -1.0, -1.0, // TL
		1.0, 1.0, 0.0, 1.0, -1.0, // TR
	}
	SQUAREIndices = []int32{
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
