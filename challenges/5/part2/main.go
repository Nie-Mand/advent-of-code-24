package main

import (
	"fmt"
	"os"
	"strings"
)

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

func solve(rules []Rule, updates [][]int) int {
	idx := index()
	idx.indexRules(rules)

	result := 0
	for _, update := range updates {
		if !idx.rightOrder(update) {
			result += middle(idx.correct(update))
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
