package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		assert.Equal(t, []int{7, 6, 4, 2, 1}, getLevels("7 6 4 2 1"), "couldn't extract levels")

		input := "test.input"
		data, err := loadInput(input)
		assert.Nil(t, err)
		require.Nil(t, err)
		assert.Equal(t, 6, len(data), "couldn't get correct reports length")
		assert.Equal(t, [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}, data, "couldn't extract input from file correctly")

		assert.Equal(t, true, hasCorrectDifference([]int{7, 6, 4, 2, 1}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, hasCorrectDifference([]int{1, 2, 7, 8, 9}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, hasCorrectDifference([]int{8, 6, 4, 4, 1}), "couldn't get correct response for isSafe")

		assert.Equal(t, true, isOneDirection([]int{9, 7, 6, 2, 1}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, isOneDirection([]int{1, 3, 2, 4, 5}), "couldn't get correct response for isSafe")
		assert.Equal(t, true, isOneDirection([]int{1, 3, 6, 7, 9}), "couldn't get correct response for isSafe")

		assert.Equal(t, true, isSafe([]int{7, 6, 4, 2, 1}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, isSafe([]int{1, 2, 7, 8, 9}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, isSafe([]int{9, 7, 6, 2, 1}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, isSafe([]int{1, 3, 2, 4, 5}), "couldn't get correct response for isSafe")
		assert.Equal(t, false, isSafe([]int{8, 6, 4, 4, 1}), "couldn't get correct response for isSafe")
		assert.Equal(t, true, isSafe([]int{1, 3, 6, 7, 9}), "couldn't get correct response for isSafe")

		assert.Equal(t, 2, safeReports(data), "couldn't get correct count of safe reports")
	})
}
