package solution

func leastBricks(wall [][]int) int {
	m := make(map[int]int, 0)
	for _, layer := range wall {
		pos := 0
		brickNums := len(layer)
		for i, brick := range layer {
			if i != brickNums-1 {
				pos = pos + brick
				m[pos]++
			}
		}
	}
	height := len(wall)
	h := 0
	for _, v := range m {
		if v > h {
			h = v
		}
	}
	return height - h
}
