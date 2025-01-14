package main

import (
	"fmt"
	"os"

	"github.com/hamillka/vk-task-2025/internal/maze"
)

func main() {
	mazeData, start, end, err := maze.ReadInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	path, err := maze.FindShortestPath(mazeData, start, end)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, p := range path {
		fmt.Printf("%d %d\n", p.Row, p.Col)
	}
	fmt.Println(".")
}
