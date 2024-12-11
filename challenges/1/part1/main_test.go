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
		assert.Nil(t, err)
		require.Nil(t, err)

		assert.NotEmpty(t, input)
		l1, l2, err := getListsFromInputSorted(input)
		assert.Nil(t, err)
		require.Nil(t, err)

		assert.Equal(t, 6, len(l1))
		assert.Equal(t, 6, len(l2))
		assert.Equal(t, []int{1, 2, 3, 3, 3, 4}, l1)
		assert.Equal(t, []int{3, 3, 3, 4, 5, 9}, l2)

		distances := calculateDistances(l1, l2)
		assert.Equal(t, 6, len(distances))
		assert.Equal(t, []int{2, 1, 0, 1, 2, 5}, distances)

		assert.Equal(t, 11, sum(distances))
	})
}
