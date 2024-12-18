package main

func step(direction string, _p P) P {
	switch direction {
	case "^":
		return p(_p.x-1, _p.y)
	case ">":
		return p(_p.x, _p.y+1)
	case "<":
		return p(_p.x, _p.y-1)
	case "v":
		return p(_p.x+1, _p.y)
	}

	return P{}
}

func after(direction string) string {
	switch direction {
	case "^":
		return ">"
	case ">":
		return "v"
	case "v":
		return "<"
	case "<":
		return "^"
	}

	return ""
}

func empty(row string, y int) string {
	return replace(row, y, ".")
}

func replace(row string, y int, with string) string {
	return row[:y] + with + row[y+1:]
}

func isGuard(cell byte) bool {
	return cell == '^' || cell == 'v' || cell == '<' || cell == '>'
}
