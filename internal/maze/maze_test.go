package maze

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type TestCase struct {
	name     string
	input    string
	expected []Point
	wantErr  bool
}

// parseInput parses test input string into maze and points
func parseInput(input string) (Maze, Point, Point, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) < 3 {
		return nil, Point{}, Point{}, fmt.Errorf("insufficient input data")
	}

	dimensions := strings.Fields(lines[0])
	if len(dimensions) != 2 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid dimensions format")
	}
	a, err := strconv.Atoi(dimensions[0])
	if err != nil || a <= 0 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid row count")
	}
	b, err := strconv.Atoi(dimensions[1])
	if err != nil || b <= 0 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid column count")
	}

	if len(lines) < a+2 {
		return nil, Point{}, Point{}, fmt.Errorf("insufficient matrix rows")
	}
	labyrinth := make([][]int, a)
	for i := 0; i < a; i++ {
		row := strings.Fields(lines[i+1])
		if len(row) != b {
			return nil, Point{}, Point{}, fmt.Errorf("invalid row length at line %d", i+2)
		}
		labyrinth[i] = make([]int, b)
		for j, val := range row {
			labyrinth[i][j], err = strconv.Atoi(val)
			if err != nil || labyrinth[i][j] < 0 || labyrinth[i][j] > 9 {
				return nil, Point{}, Point{}, fmt.Errorf("invalid cell value at row %d, column %d", i+1, j+1)
			}
		}
	}

	coords := strings.Fields(lines[a+1])
	if len(coords) != 4 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid start and finish coordinates format")
	}
	start := Point{}
	finish := Point{}
	start.Row, err = strconv.Atoi(coords[0])
	if err != nil || start.Row < 0 || start.Row >= a {
		return nil, Point{}, Point{}, fmt.Errorf("invalid start row")
	}
	start.Col, err = strconv.Atoi(coords[1])
	if err != nil || start.Col < 0 || start.Col >= b {
		return nil, Point{}, Point{}, fmt.Errorf("invalid start column")
	}
	finish.Row, err = strconv.Atoi(coords[2])
	if err != nil || finish.Row < 0 || finish.Row >= a {
		return nil, Point{}, Point{}, fmt.Errorf("invalid finish row")
	}
	finish.Col, err = strconv.Atoi(coords[3])
	if err != nil || finish.Col < 0 || finish.Col >= b {
		return nil, Point{}, Point{}, fmt.Errorf("invalid finish column")
	}

	return labyrinth, start, finish, nil
}

func TestFindShortestPath(t *testing.T) {
	tests := []TestCase{
		{
			name: "Basic 3x3 maze",
			input: `3 3
1 2 3
4 5 6
7 8 9
0 0 2 2`,
			expected: []Point{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}},
			wantErr:  false,
		},
		{
			name: "Simple path with wall",
			input: `3 3
1 0 1
1 0 1
1 1 1
0 0 2 2`,
			expected: []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}},
			wantErr:  false,
		},
		{
			name: "Multiple equal paths",
			input: `4 4
1 1 1 1
1 1 1 1
1 1 1 1
1 1 1 1
0 0 3 3`,
			expected: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 3}, {3, 3}},
			wantErr:  false,
		},
		{
			name: "Path blocked by walls",
			input: `3 3
1 1 1
0 0 0
1 1 1
0 0 2 2`,
			wantErr: true,
		},
		{
			name: "Start point is wall",
			input: `3 3
0 1 1
1 1 1
1 1 1
0 0 2 2`,
			wantErr: true,
		},
		{
			name: "End point is wall",
			input: `3 3
1 1 1
1 1 1
1 1 0
0 0 2 2`,
			wantErr: true,
		},
		{
			name: "Large maze with tricky path",
			input: `5 5
1 9 9 9 9
1 9 0 0 9
1 1 1 9 9
0 0 1 9 9
9 9 1 1 1
0 0 4 4`,
			expected: []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {3, 2}, {4, 2}, {4, 3}, {4, 4}},
			wantErr:  false,
		},
		{
			name: "Path with different weights",
			input: `4 4
1 5 5 5
1 5 5 5
1 1 1 5
5 5 1 1
0 0 3 3`,
			expected: []Point{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}, {3, 2}, {3, 3}},
			wantErr:  false,
		},
		{
			name: "Maze border test",
			input: `4 4
1 1 1 1
1 0 0 1
1 0 0 1
1 1 1 1
0 0 3 3`,
			expected: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 3}, {3, 3}},
			wantErr:  false,
		},
		{
			name: "Invalid dimensions",
			input: `0 3
1 1 1
1 1 1
1 1 1
0 0 2 2`,
			wantErr: true,
		},
		{
			name: "Invalid coordinates",
			input: `3 3
1 1 1
1 1 1
1 1 1
0 0 3 3`,
			wantErr: true,
		},
		{
			name: "Single cell path",
			input: `1 1
5
0 0 0 0`,
			expected: []Point{{0, 0}},
			wantErr:  false,
		},
		{
			name: "Spiral maze",
			input: `5 5
1 1 1 1 1
0 0 0 1 1
1 1 1 1 1
1 0 0 0 1
1 1 1 1 1
0 0 4 4`,
			expected: []Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}},
			wantErr:  false,
		},
		{
			name: "Maze with multiple weight variations",
			input: `4 4
1 9 9 1
1 1 9 1
9 1 1 1
9 9 9 1
0 0 3 3`,
			expected: []Point{{0, 0}, {1, 0}, {1, 1}, {2, 1}, {2, 2}, {2, 3}, {3, 3}},
			wantErr:  false,
		},
		{
			name: "Path through high-weight cells",
			input: `3 3
9 9 9
1 0 9
1 1 1
0 2 2 2`,
			expected: []Point{{0, 2}, {1, 2}, {2, 2}},
			wantErr:  false,
		},
		{
			name: "No path",
			input: `3 3
1 0 1
1 0 1
1 0 1
0 0 2 2`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maze, start, end, err := parseInput(tt.input)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("parseInput() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}

			path, err := FindShortestPath(maze, start, end)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindShortestPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				return
			}

			if len(path) != len(tt.expected) {
				t.Errorf("FindShortestPath() path length = %v, want %v", len(path), len(tt.expected))
				return
			}

			for i := 1; i < len(path); i++ {
				rowDiff := abs(path[i].Row - path[i-1].Row)
				colDiff := abs(path[i].Col - path[i-1].Col)
				if rowDiff+colDiff != 1 {
					t.Errorf("Invalid path: non-adjacent points %v and %v", path[i-1], path[i])
				}
			}

			if path[0] != start {
				t.Errorf("Path starts at %v, want %v", path[0], start)
			}
			if path[len(path)-1] != end {
				t.Errorf("Path ends at %v, want %v", path[len(path)-1], end)
			}

			if !reflect.DeepEqual(path, tt.expected) {
				actualWeight := calculatePathWeight(maze, path)
				expectedWeight := calculatePathWeight(maze, tt.expected)

				if actualWeight != expectedWeight {
					t.Errorf("Path weight = %d, want %d\nGot path: %v\nWant path: %v",
						actualWeight, expectedWeight, path, tt.expected)
				}
			}
		})
	}
}

// calculatePathWeight calculates total weight of the path
func calculatePathWeight(maze Maze, path []Point) int {
	weight := 0
	for _, p := range path {
		weight += maze[p.Row][p.Col]
	}
	return weight
}
