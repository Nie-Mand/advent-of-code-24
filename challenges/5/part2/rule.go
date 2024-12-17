package main

type Rule struct {
	first, next int
}

func r(first, next int) Rule {
	return Rule{first: first, next: next}
}
