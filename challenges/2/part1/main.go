package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func getLevels(input string) []int {
	items := strings.Split(input, " ")
	levels := make([]int, 0, len(items))
	num := 0
	for _, item := range items {
		fmt.Sscanf(item, "%d", &num)
		levels = append(levels, num)
	}

	return levels
}

func loadInput(filename string) ([][]int, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	data := strings.Split(strings.Trim(string(bytes), "\n"), "\n")
	reports := make([][]int, 0, len(data))
	for _, item := range data {
		reports = append(reports, getLevels(item))
	}

	return reports, nil
}

func hasCorrectDifference(items []int) bool {
	if len(items) < 1 {
		return true
	}

	for i := 1; i < len(items); i++ {
		if math.Abs(float64(items[i]-items[i-1])) < 1 {
			return false
		} else if math.Abs(float64(items[i]-items[i-1])) > 3 {
			return false
		}
	}

	return true
}

func isOneDirection(items []int) bool {
	if len(items) < 1 {
		return true
	}

	target := items[0] < items[1]

	for i := 1; i < len(items); i++ {
		if target && items[i-1] > items[i] {
			return false
		} else if !target && items[i-1] < items[i] {
			return false
		}
	}

	return true
}

func isSafe(items []int) bool {
	return isOneDirection(items) && hasCorrectDifference(items)
}

func safeReports(reports [][]int) int {
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count += 1
		}
	}

	return count
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	input := "puzzle.input"
	data, err := loadInput(input)
	handleError(err)
	fmt.Println(safeReports(data))
}
