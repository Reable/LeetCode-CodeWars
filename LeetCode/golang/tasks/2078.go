package tasks

// from left to right
func MaxDistance2078(colors []int) int {
	if len(colors) < 2 {
		return 0
	}

	temp := make(map[int]map[int]int)

	p1 := 0
	p2 := 1

	for p1 < len(colors)-1 && p1 < p2 {
		for p2 < len(colors) {
			if colors[p1] == colors[p2] || p1 == p2 {
				p2++
				continue
			}
			if object, hasObject := temp[colors[p1]]; hasObject {
				if _, hasColor := object[colors[p2]]; hasColor {
					if sum := p2 - p1; sum > object[colors[p2]] {
						object[colors[p2]] = sum
					}
				} else {
					object[colors[p2]] = p2
				}
			} else {
				_, has := temp[colors[p1]]

				if !has {
					temp[colors[p1]] = make(map[int]int)
				}

				temp[colors[p1]][colors[p2]] = p2 - p1
			}
			p2++

		}
		p1++
		p2 = p1 + 1
	}

	var result int

	for _, object := range temp {
		for _, v := range object {
			if v > result {
				result = v
			}
		}
	}

	return result
}

// from right to left
func maxDistance2078(colors []int) int {
	dist := 1

	for i := 0; i < len(colors); i++ {
		if i > 0 && colors[i-1] == colors[i] {
			continue
		}
		for j := len(colors) - 1; j > i; j-- {
			if colors[i] != colors[j] {
				dist = max(dist, j-i)
			}
		}
	}

	return dist
}
