package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filename := "test.input"
		rules, updates, err := loadInput(filename)
		assert.Nil(t, err, "Error loading input")
		require.Nil(t, err, "Error loading input")

		assert.Equal(t, Rule{first: 47, next: 53}, r(47, 53), "Expected 21 rules")
		assert.Equal(t, 21, len(rules), "Expected 21 rules")
		assert.Equal(t, r(47, 53), rules[0], "Expected 21 rules")
		assert.Equal(t, r(53, 13), rules[20], "Expected 21 rules")

		assert.Equal(t, 6, len(updates), "Expected 6 updates")
		assert.Equal(t, 5, len(updates[0]), "Expected 5 pages")
		assert.Equal(t, 29, updates[0][4], "Expected page 29")

		s := set()
		assert.Equal(t, false, s.has(1), "set should not have 1")
		s.add(1)
		assert.Equal(t, true, s.has(1), "set should have 1")

		idx := index()
		assert.NotNil(t, idx, "can't create index")
		assert.Equal(t, false, idx.has(47), "index should not have 47")
		idx.add(47, 53)
		idx.add(47, 11)
		idx.add(11, 2)
		idx.add(11, 2)
		idx.add(53, 10)
		assert.Equal(t, true, idx.has(47), "index should have 47")
		assert.Equal(t, true, idx[47].has(53), "set should have 53")
		assert.Equal(t, true, idx[47].has(11), "set should have 11")
		assert.Equal(t, true, idx.has(53), "index should have 47")
		assert.Equal(t, true, idx[53].has(10), "set should have 11")

		assert.Equal(t, true, idx.before(47, 11), "47 should be before 11")
		assert.Equal(t, true, idx.before(47, 15), "47 should be before 15")
		assert.Equal(t, true, idx.before(28, 47), "28 should be before 47")
		assert.Equal(t, true, idx.before(47, 4), "47 should be before 4")

		idx = index()
		assert.Equal(t, false, idx.has(53), "index should not have 53")
		idx.indexRules(rules)
		assert.Equal(t, true, idx.has(53), "index should have 53")
		assert.Equal(t, true, idx.rightOrder([]int{5, 47, 61, 53, 29}), "5,47,61,53,29 should be in right order")
		assert.Equal(t, false, idx.rightOrder([]int{75, 97, 47, 61, 53}), "75,97,47,61,53 should not be in right order")

		assert.Equal(t, 61, middle([]int{75, 47, 61, 53, 29}), "Expected 61")
		assert.Equal(t, 29, middle([]int{75, 29, 13}), "Expected 29")

		assert.Equal(t, 143, solve(rules, updates), "Expected 143")
	})
}
