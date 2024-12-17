package main

import (
	"os"
	"strings"
)

type P struct {
	x, y int
}

func p(x, y int) P {
	return P{x, y}
}

type Traversal struct {
	input []string
}

func traverse(input []string) Traversal {
	return Traversal{input}
}

func (t *Traversal) out(x, y int) bool {
	bx := len(t.input[0])
	by := len(t.input)

	return x < 0 || y < 0 || x >= bx || y >= by
}

func (t *Traversal) at(x, y int) string {
	if t.out(x, y) {
		return ""
	}

	return string(t.input[y][x])
}

func (t *Traversal) allStarts() []P {
	starts := []P{}
	for y, row := range t.input {
		for x := range row {
			if t.qualifies(x, y) {
				starts = append(starts, p(x, y))
			}
		}
	}

	return starts
}

func (t *Traversal) qualifies(x, y int) bool {
	if t.at(x, y) != "A" {
		return false
	}

	requiredPaths := []P{
		p(x+1, y+1),
		p(x-1, y-1),
		p(x+1, y-1),
		p(x-1, y+1),
	}

	for _, p := range requiredPaths {
		if t.out(p.x, p.y) {
			return false
		}
	}

	return true
}

func (t *Traversal) works(x, y int) bool {
	if !t.qualifies(x, y) {
		return false
	}

	crossings := 0
	rising := []P{p(x-1, y-1), p(x+1, y+1)}
	falling := []P{p(x-1, y+1), p(x+1, y-1)}

	if t.at(rising[0].x, rising[0].y) == "M" && t.at(rising[1].x, rising[1].y) == "S" {
		crossings++
	} else if t.at(rising[0].x, rising[0].y) == "S" && t.at(rising[1].x, rising[1].y) == "M" {
		crossings++
	}

	if t.at(falling[0].x, falling[0].y) == "M" && t.at(falling[1].x, falling[1].y) == "S" {
		crossings++
	} else if t.at(falling[0].x, falling[0].y) == "S" && t.at(falling[1].x, falling[1].y) == "M" {
		crossings++
	}

	return crossings == 2
}

func (t *Traversal) solve() int {
	starts := t.allStarts()
	count := 0
	for _, p := range starts {
		if t.works(p.x, p.y) {
			count++
		}
	}

	return count
}

func loadInput(filename string) ([]string, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(strings.Trim(string(b), "\n"), "\n"), nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "puzzle.input"
	input, err := loadInput(filename)
	handleError(err)
	t := traverse(input)
	println(t.solve())
}
