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

		miniFilename := "mini-test.input"
		miniInput, err := loadInput(miniFilename)
		assert.Nil(t, err, "can't read input")
		require.Nil(t, err, "can't read input")

		T := traverse(input)
		Tm := traverse(miniInput)
		assert.Equal(t, Tm.at(1, 1), "A")
		assert.Equal(t, Tm.at(2, 2), "S")

		assert.Equal(t, true, Tm.qualifies(1, 1), "can't qualify")
		assert.Equal(t, false, Tm.qualifies(2, 1), "can't qualify")

		assert.Equal(t, true, Tm.works(1, 1), "can't work")
		assert.Equal(t, P{0, 0}, p(0, 0), "can't create a point")
		assert.Equal(t, true, T.out(-1, 0), "can't check out of bounds")
		assert.Equal(t, true, T.out(-1, 0), "can't check out of bounds")
		assert.Equal(t, false, T.out(0, 1), "can't check out of bounds")
		assert.Equal(t, true, T.out(0, 10), "can't check out of bounds")

		assert.NotEmpty(t, len(T.allStarts()))
		assert.Equal(t, 1, Tm.solve())
		assert.Equal(t, 9, T.solve())
	})
}
