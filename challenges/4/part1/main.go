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

func (t *Traversal) leads(p1, p2 P) bool {
	if t.at(p1.x, p1.y) != "X" && t.at(p2.x, p2.y) != "M" {
		return false
	}

	stepX := p2.x - p1.x
	stepY := p2.y - p1.y

	for i := 0; i < 3; i++ {
		if t.at(p1.x+stepX, p1.y+stepY) != after(t.at(p1.x, p1.y)) {
			return false
		}

		p1 = p(p1.x+stepX, p1.y+stepY)
	}

	return true
}

func (t *Traversal) countLeads(p P) int {
	paths := t.paths(p.x, p.y)
	count := 0
	for _, p2 := range paths {
		if t.leads(p, p2) {
			count++
		}
	}

	return count
}

func (t *Traversal) allStarts() []P {
	starts := []P{}
	for y, row := range t.input {
		for x, c := range row {
			if c == 'X' {
				starts = append(starts, p(x, y))
			}
		}
	}

	return starts
}

func (t *Traversal) solve() int {
	starts := t.allStarts()
	count := 0
	for _, p := range starts {
		count += t.countLeads(p)
	}

	return count
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
	filename := "puzzle.input"
	input, err := loadInput(filename)
	handleError(err)
	t := traverse(input)
	println(t.solve())
}
