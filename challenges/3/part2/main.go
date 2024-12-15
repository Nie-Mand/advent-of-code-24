package main

import (
	"fmt"
	"os"
	"strings"
)

type Op struct {
	is   string
	l, r int
}

func loadInput(filename string) (string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return strings.Trim(string(file), "\n"), nil
}

func cut(s string, start, size int) string {
	if len(s) < start+size {
		return s
	}

	return s[start : start+size]
}

func mul(a, b int) Op {
	return Op{"mul", a, b}
}

func do() Op {
	return Op{"do", 0, 0}
}

func dont() Op {
	return Op{"dont", 0, 0}
}

func isOpNil(op Op) bool {
	return op.is == "mul" && op.l == 0 && op.r == 0
}

func toOperation(s string) Op {
	out := mul(0, 0)
	_, err := fmt.Sscanf(s, "mul(%d,%d)", &out.l, &out.r)
	if err != nil {
		return mul(0, 0)
	}

	return out
}

func getNextOperation(s string) (Op, string) {
	for i := 0; i < len(s); i++ {
		if cut(s, i, 7) == "don't()" {
			return dont(), s[i+7:]
		} else if cut(s, i, 4) == "do()" {
			return do(), s[i+4:]
		} else if cut(s, i, 4) == "mul(" {
			op := cut(s, i, 4)
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

func extractOperations(s string) []Op {
	out := []Op{}
	var op Op
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
	enabled := true
	for _, op := range ops {
		if op.is == "do" {
			enabled = true
		} else if op.is == "dont" {
			enabled = false
		} else if enabled {
			result += op.l * op.r
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
	in, err := loadInput(filename)
	handleError(err)

	fmt.Println(solve(in))
}
