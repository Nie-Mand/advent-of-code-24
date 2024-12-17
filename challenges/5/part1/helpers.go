package main

func middle(pages []int) int {
	if len(pages)%2 == 0 {
		return pages[len(pages)/2-1]
	}

	return pages[len(pages)/2]
}
