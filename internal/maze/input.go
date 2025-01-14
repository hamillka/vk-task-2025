package maze

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput reads and validates input from stdin
func ReadInput() (Maze, Point, Point, error) {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return nil, Point{}, Point{}, fmt.Errorf("failed to read maze dimensions")
	}
	dims := strings.Fields(scanner.Text())
	if len(dims) != 2 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid maze dimensions format")
	}

	rows, err1 := strconv.Atoi(dims[0])
	cols, err2 := strconv.Atoi(dims[1])
	if err1 != nil || err2 != nil || rows <= 0 || cols <= 0 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid maze dimensions")
	}

	maze := make(Maze, rows)
	for i := 0; i < rows; i++ {
		if !scanner.Scan() {
			return nil, Point{}, Point{}, fmt.Errorf("failed to read maze row")
		}

		numbers := strings.Fields(scanner.Text())
		if len(numbers) != cols {
			return nil, Point{}, Point{}, fmt.Errorf("invalid maze row length")
		}

		maze[i] = make([]int, cols)
		for j, num := range numbers {
			val, err := strconv.Atoi(num)
			if err != nil || val < 0 || val > 9 {
				return nil, Point{}, Point{}, fmt.Errorf("invalid maze cell value")
			}
			maze[i][j] = val
		}
	}

	if !scanner.Scan() {
		return nil, Point{}, Point{}, fmt.Errorf("failed to read start/end points")
	}
	coords := strings.Fields(scanner.Text())
	if len(coords) != 4 {
		return nil, Point{}, Point{}, fmt.Errorf("invalid start/end points format")
	}

	start := Point{}
	end := Point{}
	start.Row, err1 = strconv.Atoi(coords[0])
	start.Col, err2 = strconv.Atoi(coords[1])
	if err1 != nil || err2 != nil {
		return nil, Point{}, Point{}, fmt.Errorf("invalid start point")
	}

	end.Row, err1 = strconv.Atoi(coords[2])
	end.Col, err2 = strconv.Atoi(coords[3])
	if err1 != nil || err2 != nil {
		return nil, Point{}, Point{}, fmt.Errorf("invalid end point")
	}

	return maze, start, end, nil
}
