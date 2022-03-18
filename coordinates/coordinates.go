package coordinates

// Coord is a Cartesian coordinate
type Coord struct {
	X, Y, Z int
}

// Transform moves a coordinate by the vector "v"
func (c *Coord) Transform(v Coord) {
	c.X += v.X
	c.Y += v.Y
	c.Z += v.Z
}

// RotateX90 rotates a coordinate 90 degrees around the x axis
func (c *Coord) RotateX90() {
	c.Y, c.Z = -c.Z, c.Y
}

// RotateY90 rotates a coordinate 90 degrees around the y axis
func (c *Coord) RotateY90() {
	c.X, c.Z = c.Z, -c.X
}

// RotateZ90 rotates a coordinate 90 degrees around the z axis
func (c *Coord) RotateZ90() {
	c.X, c.Y = -c.Y, c.X
}

// ManhattanDistance returns the Manhattan distance between two coordinates
func ManhattanDistance(a, b Coord) int {
	d := 0

	if b.X > a.X {
		d += b.X - a.X
	} else {
		d += a.X - b.X
	}

	if b.Y > a.Y {
		d += b.Y - a.Y
	} else {
		d += a.Y - b.Y
	}

	if b.Z > a.Z {
		d += b.Z - a.Z
	} else {
		d += a.Z - b.Z
	}

	return d
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
		if c.Z < minCoord.Z {
			minCoord.Z = c.Z
		}

		// Check for larger values
		if c.X > maxCoord.X {
			maxCoord.X = c.X
		}
		if c.Y > maxCoord.Y {
			maxCoord.Y = c.Y
		}
		if c.Z > maxCoord.Z {
			maxCoord.Z = c.Z
		}
	}

	return minCoord, maxCoord
}
