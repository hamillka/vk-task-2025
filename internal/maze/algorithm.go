package maze

import (
	"container/heap"
	"fmt"
	"math"

	"github.com/hamillka/vk-task-2025/internal/priorityqueue"
)

// manhattanDistance calculates the Manhattan distance between two points
func manhattanDistance(p1, p2 Point) int {
	return abs(p1.Row-p2.Row) + abs(p1.Col-p2.Col)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// FindShortestPath finds the shortest path using A* algorithm
func FindShortestPath(maze Maze, start, end Point) ([]Point, error) {
	if !maze.IsValidPoint(start) {
		return nil, fmt.Errorf("invalid start point")
	}
	if !maze.IsValidPoint(end) {
		return nil, fmt.Errorf("invalid end point")
	}

	rows, cols := len(maze), len(maze[0])
	gScores := make([][]int, rows)
	previous := make([][]Point, rows)
	inOpenSet := make([][]bool, rows)

	for i := range gScores {
		gScores[i] = make([]int, cols)
		previous[i] = make([]Point, cols)
		inOpenSet[i] = make([]bool, cols)
		for j := range gScores[i] {
			gScores[i][j] = math.MaxInt
		}
	}

	gScores[start.Row][start.Col] = maze[start.Row][start.Col]
	startHeuristic := manhattanDistance(start, end)

	pq := make(priorityqueue.PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &priorityqueue.Item{
		Value:    start,
		Priority: gScores[start.Row][start.Col] + startHeuristic,
		GScore:   gScores[start.Row][start.Col],
	})
	inOpenSet[start.Row][start.Col] = true

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*priorityqueue.Item)
		currentPoint := current.Value.(Point)
		inOpenSet[currentPoint.Row][currentPoint.Col] = false

		if currentPoint == end {
			break
		}

		if current.GScore > gScores[currentPoint.Row][currentPoint.Col] {
			continue
		}

		for _, neighbor := range maze.GetNeighbors(currentPoint) {
			tentativeGScore := gScores[currentPoint.Row][currentPoint.Col] + maze[neighbor.Row][neighbor.Col]

			if tentativeGScore < gScores[neighbor.Row][neighbor.Col] {
				previous[neighbor.Row][neighbor.Col] = currentPoint
				gScores[neighbor.Row][neighbor.Col] = tentativeGScore

				if !inOpenSet[neighbor.Row][neighbor.Col] {
					hScore := manhattanDistance(neighbor, end)
					heap.Push(&pq, &priorityqueue.Item{
						Value:    neighbor,
						Priority: tentativeGScore + hScore,
						GScore:   tentativeGScore,
					})
					inOpenSet[neighbor.Row][neighbor.Col] = true
				}
			}
		}
	}

	if gScores[end.Row][end.Col] == math.MaxInt {
		return nil, fmt.Errorf("no path found")
	}

	path := make([]Point, 0)
	current := end
	for current != start {
		path = append([]Point{current}, path...)
		current = previous[current.Row][current.Col]
	}
	path = append([]Point{start}, path...)

	return path, nil
}
