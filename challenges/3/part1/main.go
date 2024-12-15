package main

import (
	"fmt"
	"os"
	"strings"
)

func loadInput(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(file), "\n"), nil
}

func isOpNil(op []int) bool {
	return op[0] == 0 && op[1] == 0
}

func toOperation(s string) []int {
	out := make([]int, 2)
	_, err := fmt.Sscanf(s, "mul(%d,%d)", &out[0], &out[1])
	if err != nil {
		out[0] = 0
		out[1] = 0
	}

	return out
}

func getNextOperation(s string) ([]int, string) {
	if len(s) < 8 {
		return []int{0, 0}, ""
	}

	for i := 0; i < len(s)-7; i++ {
		if s[i:i+4] == "mul(" {
			op := s[i : i+4]
			for j := i + 4; j < len(s); j++ {
				if s[j] == 'm' {
					break
				}
				if s[j] == ')' {
					return toOperation(op + s[i+4:j+1]), s[j+1:]
				}
			}
		}

	}

	return toOperation(s), ""
}

func extractOperations(s string) [][]int {
	out := [][]int{}
	var op []int
	for s != "" {
		op, s = getNextOperation(s)
		if !isOpNil(op) {
			out = append(out, op)
		}
	}

	return out
}

func solve(s string) int {
	ops := extractOperations(s)
	result := 0
	for _, op := range ops {
		result += op[0] * op[1]
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
	in, err := loadInput(filename)
	handleError(err)

	fmt.Println(solve(in))
}
