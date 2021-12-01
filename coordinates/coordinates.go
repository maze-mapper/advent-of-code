package coordinates

// Coord is a Cartesian coordinate
type Coord struct {
	X, Y int
}

// Range returns two coordinates representing the minimum and maximum values in each dimension for a set of coordinates
func Range(coords []Coord) (Coord, Coord) {
	// Set min and max coordinates to initially be the first coordinate
	minCoord := coords[0]
	maxCoord := coords[0]

	for _, c := range coords[1:] {
		// Check for smaller values
		if c.X < minCoord.X {
                        minCoord.X = c.X
                }
		if c.Y < minCoord.Y {
			minCoord.Y = c.Y
		}

		// Check for larger values
		if c.X > maxCoord.X {
			maxCoord.X = c.X
		}
		if c.Y > maxCoord.Y {
                        maxCoord.Y = c.Y
                }
	}

	return minCoord, maxCoord
}

