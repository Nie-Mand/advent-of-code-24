package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(content), "\n"), nil
}

func getListsFromInput(input string) ([]int, []int, error) {
	lines := strings.Split(input, "\n")
	len := len(lines)
	l1 := make([]int, len)
	l2 := make([]int, len)

	for i := 0; i < len; i++ {
		fmt.Sscanf(lines[i], "%d %d", &l1[i], &l2[i])
	}

	return l1, l2, nil
}

func _getOccurances(l []int) map[int]int {
	occurances := make(map[int]int)
	for i := 0; i < len(l); i++ {
		key := l[i]
		_, ok := occurances[key]
		if ok {
			occurances[key] += 1
		} else {
			occurances[key] = 1
		}
	}

	return occurances
}

func calculateSimilarities(l1, l2 []int) []int {
	occurances := _getOccurances(l2)

	d := make([]int, len(l1))
	for i := 0; i < len(l1); i++ {
		d[i] = l1[i] * occurances[l1[i]]
	}
	return d
}

func sum(d []int) int {
	_sum := 0
	for _, item := range d {
		_sum += item
	}

	return _sum
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputName := "puzzle.input"
	input, err := getInput(inputName)
	handleError(err)

	l1, l2, err := getListsFromInput(input)
	handleError(err)

	distances := calculateSimilarities(l1, l2)
	fmt.Println("totalSimilarities", sum(distances))
}
