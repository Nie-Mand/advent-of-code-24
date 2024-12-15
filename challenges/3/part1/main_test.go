package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filename := "test.input"
		in, err := loadInput(filename)
		assert.Nil(t, err, "Can't load input")
		require.Nil(t, err, "Can't load input")
		assert.Equal(t, "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", in, "Can't load input")

		assert.Equal(t, true, isOpNil([]int{0, 0}), "Can't check if operation is nil")
		assert.Equal(t, false, isOpNil([]int{1, 0}), "Can't check if operation is nil")
		assert.Equal(t, false, isOpNil([]int{0, 1}), "Can't check if operation is nil")
		assert.Equal(t, false, isOpNil([]int{1, 1}), "Can't check if operation is nil")

		assert.Equal(t, []int{4, 3}, toOperation("mul(4,3)"), "Can't convert string to operation")
		assert.Equal(t, []int{3, 4}, toOperation("mul(3,4)"), "Can't convert string to operation")
		assert.Equal(t, []int{0, 0}, toOperation("mul(4*"), "Can't convert string to operation")
		assert.Equal(t, []int{0, 0}, toOperation("mul(6,9!"), "Can't convert string to operation")
		assert.Equal(t, []int{0, 0}, toOperation("?(12,34)"), "Can't convert string to operation")
		assert.Equal(t, []int{0, 0}, toOperation("mul ( 2 , 4 )"), "Can't convert string to operation")

		ops := []struct {
			in   string
			op   []int
			next string
		}{
			{"mul(4,", []int{0, 0}, ""},
			{"mul(4,   mul", []int{0, 0}, ""},
			{"xyzmul(4,3)", []int{4, 3}, ""},
			{"xyzmul(4,3)   ", []int{4, 3}, "   "},
			{"mul(4,3)xmul(3,4)", []int{4, 3}, "xmul(3,4)"},
			{"xyzmul(4,3) mul(1,", []int{4, 3}, " mul(1,"},
			{"mul(32,64]then(mul(11,8)", []int{11, 8}, ""},
		}
		for _, item := range ops {
			op, next := getNextOperation(item.in)
			assert.Equal(t, item.op, op, "Can't get next operation for %s", item.in)
			assert.Equal(t, item.next, next, "Can't get remaining string for %s", item.in)
		}

		fullExtractionOps := []struct {
			in  string
			out [][]int
		}{
			{"mul(4,", [][]int{}},
			{"mul(4,   mul", [][]int{}},
			{"xyzmul(4,3)", [][]int{{4, 3}}},
			{"xyzmul(4,3)   ", [][]int{{4, 3}}},
			{"mul(4,3)xmul(3,4)", [][]int{{4, 3}, {3, 4}}},
			{"xyzmul(4,3) mul(1,", [][]int{{4, 3}}},
			{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", [][]int{{2, 4}, {5, 5}, {11, 8}, {8, 5}}},
		}
		for _, item := range fullExtractionOps {
			assert.Equal(t, item.out, extractOperations(item.in), "Can't extract all operation for %s", item.in)
		}

		assert.Equal(t, 161, solve("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"), "Can't solve the problem")
	})
}
