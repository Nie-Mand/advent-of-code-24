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

func (t *Traversal) next(x, y int) []P {
	possible := t.paths(x, y)
	next := []P{}
	for _, p := range possible {
		if after(t.at(x, y)) == t.at(p.x, p.y) {
			next = append(next, p)
		}
	}

	return next
}

func (t *Traversal) paths(x, y int) []P {
	paths := []P{
		p(x+1, y),
		p(x+1, y+1),
		p(x, y+1),
		p(x-1, y+1),
		p(x-1, y),
		p(x-1, y-1),
		p(x, y-1),
		p(x+1, y-1),
	}

	possible := []P{}
	for _, p := range paths {
		if !t.out(p.x, p.y) {
			possible = append(possible, p)
		}
	}

	return possible
}

func after(c string) string {
	switch c {
	case "X":
		return "M"
	case "M":
		return "A"
	case "A":
		return "S"
	}
	return ""
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

}
