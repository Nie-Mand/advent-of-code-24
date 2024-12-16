package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filename := "test.input"
		input, err := loadInput(filename)
		assert.Nil(t, err, "can't read input")
		require.Nil(t, err, "can't read input")
		assert.Equal(t, 10, len(input), "can't read all lines")
		assert.Equal(t, 10, len(input[0]), "can't read all columns")

		assert.Equal(t, "M", after("X"), "can't get the next char")
		assert.Equal(t, "A", after("M"), "can't get the next char")
		assert.Equal(t, "S", after("A"), "can't get the next char")
		assert.Equal(t, "", after("S"), "can't get the next char")

		T := traverse(input)
		assert.Equal(t, T.at(3, 0), "S")
		assert.Equal(t, T.at(3, 1), "M")
		assert.Equal(t, P{0, 0}, p(0, 0), "can't create a point")
		assert.Equal(t, true, T.out(-1, 0), "can't check out of bounds")
		assert.Equal(t, true, T.out(-1, 0), "can't check out of bounds")
		assert.Equal(t, false, T.out(0, 1), "can't check out of bounds")
		assert.Equal(t, true, T.out(0, 10), "can't check out of bounds")
		assert.Equal(t, []P{{1, 0}, {1, 1}, {0, 1}}, T.paths(0, 0), "can't get all paths")

		assert.Equal(t, []P{{5, 1}, {3, 1}}, T.next(4, 0))
		assert.Equal(t, []P{{6, 0}, {5, 1}}, T.next(5, 0))

		assert.Equal(t, true, T.leads(P{4, 0}, P{5, 1}))
		assert.Equal(t, false, T.leads(P{4, 0}, P{3, 1}))
		assert.Equal(t, 1, T.countLeads(P{4, 0}))

		assert.Equal(t, 19, len(T.allStarts()))
		assert.Equal(t, 18, T.solve())
	})
}
