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
		if evaluateSafetyWithRetries(report, 1) {
			count += 1
		}
	}

	return count
}

func dropItem(data []int, idx int) []int {
	copy := make([]int, 0, len(data)-1)
	for i, item := range data {
		if i == idx {
			continue
		}

		copy = append(copy, item)
	}

	return copy
}

func evaluateSafetyAt(data []int, idx int) bool {
	return isSafe(dropItem(data, idx))
}

func evaluateSafetyWithRetries(data []int, retries int) bool {
	if isSafe(data) {
		return true
	}

	fixed := false
	for idx := 0; idx < len(data); idx++ {
		if isSafe(dropItem(data, idx)) {
			fixed = true
			retries -= 1
			if retries == 0 {
				return true
			}
		}
	}

	return fixed && retries == 0
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
