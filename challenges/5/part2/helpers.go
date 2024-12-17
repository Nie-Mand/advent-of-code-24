package main

func swap(pages []int, i int) []int {
	pages[i], pages[i+1] = pages[i+1], pages[i]
	return pages
}

func middle(pages []int) int {
	if len(pages)%2 == 0 {
		return pages[len(pages)/2-1]
	}

	return pages[len(pages)/2]
}
