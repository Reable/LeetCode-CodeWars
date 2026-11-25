package tasks

// 2,2,1,1,1,2,2
func MajorityElement169(nums []int) int {
	candidate, count := 0, 0

	for _, num := range nums {
		switch {
		case candidate == num && count > 0:
			count++
		case count == 0:
			candidate = num
			count = 1
		default:
			count--
		}
	}

	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}
	return -1
}
