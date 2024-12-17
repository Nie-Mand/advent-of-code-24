package main

import (
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	first, next int
}

func r(first, next int) Rule {
	return Rule{first: first, next: next}
}

type Set map[int]bool

func set() Set {
	return Set{}
}

func (s Set) add(i int) {
	s[i] = true
}

func (s Set) has(i int) bool {
	_, ok := s[i]
	return ok
}

func loadInput(filename string) ([]Rule, [][]int, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Rule{}, [][]int{}, err
	}

	data := strings.Split(string(b), "\n")

	rules := make([]Rule, 0, len(data))
	updates := make([][]int, 0, len(data))

	idx := 0
	phase := 0
	for idx < len(data) && phase < 2 {
		line := data[idx]

		if line == "" {
			phase += 1
			idx += 1
			continue
		}

		switch phase {
		case 0:
			rule := Rule{}
			fmt.Sscanf(line, "%d|%d\n", &rule.first, &rule.next)

			rules = append(rules, rule)
			break
		case 1:
			pages := strings.Split(line, ",")
			update := make([]int, len(pages))
			for i, page := range pages {
				fmt.Sscanf(page, "%d", &update[i])
			}

			updates = append(updates, update)
			break
		}
		idx++
	}

	return rules, updates, nil
}

type Index map[int]Set

func index() Index {
	return map[int]Set{}
}

func (idx Index) has(v int) bool {
	_, ok := idx[v]
	return ok
}

func (idx Index) add(first, second int) {
	if _, ok := idx[first]; !ok {
		idx[first] = set()
	}

	if _, ok := idx[second]; !ok {
		idx[second] = set()
	}

	idx[first].add(second)
}

func (idx Index) before(first, second int) bool {
	if !idx.has(second) || !idx.has(first) {
		return true
	}

	return idx[first].has(second)
}

func (idx Index) indexRules(rules []Rule) {
	for _, rule := range rules {
		idx.add(rule.first, rule.next)
	}
}

func (idx Index) rightOrder(pages []int) bool {
	for i := 0; i < len(pages)-1; i++ {
		if !idx.before(pages[i], pages[i+1]) {
			return false
		}
	}

	return true
}

func middle(pages []int) int {
	if len(pages)%2 == 0 {
		return pages[len(pages)/2-1]
	}

	return pages[len(pages)/2]
}

func solve(rules []Rule, updates [][]int) int {
	idx := index()
	idx.indexRules(rules)

	result := 0
	for _, update := range updates {
		if idx.rightOrder(update) {
			result += middle(update)
		}
	}

	return result
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "puzzle.input"
	rules, updates, err := loadInput(filename)
	handleError(err)

	fmt.Println(solve(rules, updates))
}
