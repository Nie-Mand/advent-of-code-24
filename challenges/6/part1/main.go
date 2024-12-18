package main

import (
	"fmt"
	"os"
	"strings"
)

type P struct {
	x, y int
}

func p(x, y int) P {
	return P{x, y}
}

func loadTest(filename string) (Grid, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return Grid{}, err
	}

	return newGrid(strings.Split(strings.Trim(string(b), "\n"), "\n")), nil
}

func solve(grid Grid) int {
	for {
		guard := grid.guardAt()

		if grid.out(guard) {
			break
		}

		grid.move()
	}

	return grid.visited()
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "puzzle.input"
	grid, err := loadTest(filename)
	handleError(err)

	fmt.Println(solve(grid))
}
