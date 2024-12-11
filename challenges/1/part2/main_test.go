package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		inputName := "test.input"
		input, err := getInput(inputName)
		require.Nil(t, err)

		l1, l2, err := getListsFromInput(input)
		require.Nil(t, err)
		assert.Equal(t, []int{3, 4, 2, 1, 3, 3}, l1)
		assert.Equal(t, []int{4, 3, 5, 3, 9, 3}, l2)

		similarities := calculateSimilarities(l1, l2)
		assert.Equal(t, 6, len(similarities))
		assert.Equal(t, []int{9, 4, 0, 0, 9, 9}, similarities)
		assert.Equal(t, 31, sum(similarities))
	})
}
