package main

type Grid struct {
	g        []string
	_visited [][]bool
}

func newGrid(g []string) Grid {
	grid := Grid{
		g: g,
	}

	grid._visited = make([][]bool, len(g))
	for i := range grid._visited {
		grid._visited[i] = make([]bool, len(g))
	}

	return grid
}

func (grid Grid) at(i int) string {
	return grid.g[i]
}

func (grid Grid) visited() int {
	count := 0
	for _, row := range grid._visited {
		for _, v := range row {
			if v {
				count += 1
			}
		}
	}

	return count
}

func (grid Grid) p(x, y int) byte {
	return grid.g[x][y]
}

func (grid Grid) guardAt() P {
	for x, row := range grid.g {
		for y, _ := range row {
			if isGuard(grid.p(x, y)) {
				grid._visited[x][y] = true
				return P{x, y}
			}
		}
	}

	return P{-1, -1}
}

func (grid Grid) out(_p P) bool {
	l := len(grid.g)
	return !(_p.x >= 0 && _p.y >= 0 && _p.x < l && _p.y < l)
}

func (grid Grid) move() {
	guard := grid.guardAt()
	oldGuard := guard
	direction := string(grid.g[guard.x][guard.y])

	for !grid.out(guard) {
		next := step(string(grid.g[oldGuard.x][oldGuard.y]), guard)
		if grid.out(next) {
			guard = next
			break
		}

		if grid.g[next.x][next.y] == '#' {
			break
		}

		guard = next
		grid._visited[guard.x][guard.y] = true
	}

	grid.g[oldGuard.x] = empty(grid.g[oldGuard.x], oldGuard.y)

	if !grid.out(guard) {
		grid.g[guard.x] = replace(grid.g[guard.x], guard.y, after(direction))
	}
}
