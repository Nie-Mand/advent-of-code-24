package main

type Index map[int]Set

func index() Index {
	return map[int]Set{}
}

func (idx Index) has(v int) bool {
	_, ok := idx[v]
	return ok
}

func (idx Index) add(first, second int) {
	if _, ok := idx[first]; !ok {
		idx[first] = set()
	}

	if _, ok := idx[second]; !ok {
		idx[second] = set()
	}

	idx[first].add(second)
}

func (idx Index) before(first, second int) bool {
	if !idx.has(second) || !idx.has(first) {
		return true
	}

	return idx[first].has(second)
}

func (idx Index) indexRules(rules []Rule) {
	for _, rule := range rules {
		idx.add(rule.first, rule.next)
	}
}

func (idx Index) rightOrder(pages []int) bool {
	for i := 0; i < len(pages)-1; i++ {
		if !idx.before(pages[i], pages[i+1]) {
			return false
		}
	}

	return true
}

func (idx Index) failsAt(pages []int) int {
	for i := 0; i < len(pages)-1; i++ {
		if !idx.before(pages[i], pages[i+1]) {
			return i
		}
	}

	return -1
}

func (idx Index) correct(pages []int) []int {
	id := idx.failsAt(pages)
	if id == -1 {
		return pages
	}

	return idx.correct(swap(pages, id))
}
