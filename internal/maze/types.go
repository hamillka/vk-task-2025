package maze

// Point represents coordinates in the maze
type Point struct {
	Row, Col int
}

// Maze represents the maze structure
type Maze [][]int

// IsValidPoint checks if the point is within maze boundaries and not a wall
func (m Maze) IsValidPoint(p Point) bool {
	return p.Row >= 0 && p.Row < len(m) &&
		p.Col >= 0 && p.Col < len(m[0]) &&
		m[p.Row][p.Col] != 0
}

// GetNeighbors returns valid neighboring points
func (m Maze) GetNeighbors(p Point) []Point {
	directions := []Point{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	neighbors := make([]Point, 0, 4)

	for _, dir := range directions {
		newPoint := Point{p.Row + dir.Row, p.Col + dir.Col}
		if m.IsValidPoint(newPoint) {
			neighbors = append(neighbors, newPoint)
		}
	}

	return neighbors
}
