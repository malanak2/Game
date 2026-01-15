package Objects

var (
	SQUAREVertices = []float32{
		// Positions 3	// Texture Coord 2
		-1.0, -1.0, 0.0, 1.0, 1.0, // BL
		1.0, -1.0, 0.0, -1.0, 1.0, // BR
		-1.0, 1.0, 0.0, 1.0, -1.0, // TL
		1.0, 1.0, 0.0, -1.0, -1.0, // TR
	}
	SQUAREIdices = []int32{
		0, 1, 2,
		1, 2, 3,
	}
)
