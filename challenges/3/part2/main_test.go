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
		assert.Nil(t, err, "Can't load input")
		require.Nil(t, err, "Can't load input")
		assert.Equal(t, "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", input, "Can't load input")

		assert.Equal(t, false, isOpNil(mul(4, 3)), "Can't check if operation is nil")
		assert.Equal(t, true, isOpNil(mul(0, 0)), "Can't check if operation is nil")
		assert.Equal(t, false, isOpNil(do()), "Can't check if operation is nil")
		assert.Equal(t, false, isOpNil(dont()), "Can't check if operation is nil")
		assert.Equal(t, "do()  ", cut("do()  ", 0, 6), "Can't cut string")

		assert.Equal(t, Op{is: "mul", l: 4, r: 3}, mul(4, 3), "Can't convert 4*3 to operation")
		assert.Equal(t, Op{is: "do"}, do(), "Can't convert done string to operation")
		assert.Equal(t, Op{is: "dont"}, dont(), "Can't convert dont string to operation")
		assert.Equal(t, mul(4, 3), toOperation("mul(4,3)"), "Can't convert string to operation")
		assert.Equal(t, mul(3, 4), toOperation("mul(3,4)"), "Can't convert string to operation")
		assert.Equal(t, mul(0, 0), toOperation("mul(4*"), "Can't convert string to operation")
		assert.Equal(t, mul(0, 0), toOperation("mul(6,9!"), "Can't convert string to operation")
		assert.Equal(t, mul(0, 0), toOperation("?(12,34)"), "Can't convert string to operation")
		assert.Equal(t, mul(0, 0), toOperation("mul ( 2 , 4 )"), "Can't convert string to operation")

		ops := []struct {
			in   string
			op   Op
			next string
		}{
			{"mul(4,", mul(0, 0), ""},
			{"do() ", do(), " "},
			{"don't()", dont(), ""},
			{"mul(4,   mul", mul(0, 0), ""},
			{"xyzmul(4,3)", mul(4, 3), ""},
			{"xyzmul(4,3)   ", mul(4, 3), "   "},
			{"xyzdo()   ", do(), "   "},
			{"xydon't() mul(4,3)44", dont(), " mul(4,3)44"},
			{"mul(4,3)xmul(3,4)", mul(4, 3), "xmul(3,4)"},
			{"xyzmul(4,3) mul(1,", mul(4, 3), " mul(1,"},
			{"mul(32,64]then(mul(11,8)", mul(11, 8), ""},
		}
		for _, item := range ops {
			op, next := getNextOperation(item.in)
			assert.Equal(t, item.op, op, "Can't get next operation for %s", item.in)
			assert.Equal(t, item.next, next, "Can't get remaining string for %s", item.in)
		}

		fullExtractionOps := []struct {
			in  string
			out []Op
		}{
			{"mul(4,", []Op{}},
			{"mul(4,   mul", []Op{}},
			{"xyzmul(4,3)don't()", []Op{mul(4, 3), dont()}},
			{"xyzmul(4,3)   ", []Op{mul(4, 3)}},
			{"mul(4,3)do()xmul(3,4)", []Op{mul(4, 3), do(), mul(3, 4)}},
			{"xyzmul(4,3) mul(1,", []Op{mul(4, 3)}},
			{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", []Op{mul(2, 4), mul(5, 5), mul(11, 8), mul(8, 5)}},
			{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", []Op{mul(2, 4), dont(), mul(5, 5), mul(11, 8), do(), mul(8, 5)}},
		}
		for _, item := range fullExtractionOps {
			assert.Equal(t, item.out, extractOperations(item.in), "Can't extract all operation for %s", item.in)
		}

		assert.Equal(t, 48, solve(input), "Can't solve the problem")
	})
}
