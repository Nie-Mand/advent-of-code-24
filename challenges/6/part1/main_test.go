package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filename := "test.input"
		grid, err := loadTest(filename)
		assert.Nil(t, err)
		require.Nil(t, err)
		assert.Equal(t, 10, len(grid.g))
		assert.Equal(t, 10, len(grid.at(0)))
		assert.Equal(t, "......#...", grid.at(9))
		assert.Equal(t, P{x: 0, y: 1}, p(0, 1))

		assert.Equal(t, p(6, 4), grid.guardAt())

		assert.Equal(t, true, isGuard(grid.p(6, 4)))
		grid.g[6] = ".#..>....."
		assert.Equal(t, true, isGuard(grid.p(6, 4)))
		assert.Equal(t, false, isGuard(grid.p(6, 6)))

		assert.Equal(t, p(6, 4), grid.guardAt())
		grid.g[6] = ".#..^....."

		assert.Equal(t, false, grid.out(p(1, 1)))
		assert.Equal(t, true, grid.out(p(-1, 1)))
		assert.Equal(t, true, grid.out(p(1, 10)))
		assert.Equal(t, false, grid.out(p(1, 9)))
		assert.Equal(t, p(5, 4), step("^", p(6, 4)))
		assert.Equal(t, p(6, 5), step(">", p(6, 4)))
		assert.Equal(t, p(6, 3), step("<", p(6, 4)))
		assert.Equal(t, p(7, 4), step("v", p(6, 4)))
		assert.Equal(t, "<", after("v"))
		assert.Equal(t, "^", after("<"))
		assert.Equal(t, ">", after("^"))
		assert.Equal(t, "v", after(">"))
		assert.Equal(t, ".........#", grid.g[1])
		assert.Equal(t, ".........#", replace("....>....#", 4, "."))
		assert.Equal(t, ".........#", empty("....>....#", 4))
		grid.move()
		assert.Equal(t, "....>....#", grid.g[1], "invalid for new guard")
		assert.Equal(t, ".#........", grid.g[6], "invalid for old guard")
		assert.Equal(t, 6, grid.visited(), "invalid number of visited")
		grid.move()
		grid.move()
		grid.move()
		assert.Equal(t, ".#^.......", grid.g[6], "invalid for new guard")
		grid.move()
		grid.move()
		assert.Equal(t, "......v#..", grid.g[4], "invalid for new guard")
		grid.move()
		assert.Equal(t, "#.....<...", grid.g[8], "invalid for new guard")
		grid.move()
		grid.move()
		grid.move()
		assert.Equal(t, ".......v#.", grid.g[7], "invalid for new guard")
		grid.move()
		assert.Equal(t, "......#...", grid.g[9], "invalid for new guard")

		assert.Equal(t, 41, solve(grid), "invalid result")
	})
}
