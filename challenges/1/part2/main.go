package main

import (
	"fmt"
	"math"
	"os"
	"slices"
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

	slices.Sort(l1)
	slices.Sort(l2)

	return l1, l2, nil
}

func calculateDistances(l1, l2 []int) []int {
	d := make([]int, len(l1))
	for i := 0; i < len(d); i++ {
		d[i] = int(math.Abs(float64(l1[i] - l2[i])))
	}
	return d
}

func totalDistances(d []int) int {
	sum := 0
	for _, item := range d {
		sum += item
	}

	return sum
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

	distances := calculateDistances(l1, l2)
	fmt.Println("totalDifferences", totalDistances(distances))
}
