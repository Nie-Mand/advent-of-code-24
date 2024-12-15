package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		input := "test.input"
		data, err := loadInput(input)
		assert.Nil(t, err)
		require.Nil(t, err)
		assert.Equal(t, 6, len(data), "couldn't get correct reports length")
		assert.Equal(t, []int{7, 6, 4, 2, 1}, data[0], "couldn't extract input from file correctly")

		assert.Equal(t, []int{7, 4, 2, 1}, dropItem(data[0], 1), "couldn't dropItem item from data correctly")
		assert.Equal(t, false, evaluateSafetyAt([]int{1, 3, 2, 4, 5}, 0), "couldn't evaluateSafetyAt from data correctly")
		assert.Equal(t, true, evaluateSafetyAt([]int{1, 3, 2, 4, 5}, 1), "couldn't evaluateSafetyAt from data correctly")
		assert.Equal(t, false, evaluateSafetyWithRetries([]int{1, 3, 2, 4, 5}, 0), "couldn't evaluateSafetyWithRetries from data correctly")
		assert.Equal(t, true, evaluateSafetyWithRetries([]int{1, 3, 2, 4, 5}, 1), "couldn't evaluateSafetyWithRetries from data correctly")

		out := []bool{true, false, false, true, true, true}
		for i, report := range data {
			assert.Equal(t, out[i], evaluateSafetyWithRetries(report, 1), "couldn't evaluate safety of report correctly")
		}

		assert.Equal(t, 4, safeReports(data), "couldn't get correct safe reports count")
	})
}
