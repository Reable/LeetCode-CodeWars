package tasks

func MajorityElement229(nums []int) []int {
	c1, c2, count1, count2 := 0, 0, 0, 0
	for _, num := range nums {
		switch {
		case num == c1 && count1 > 0:
			count1++
		case num == c2 && count2 > 0:
			count2++
		case count1 == 0:
			c1 = num
			count1++
		case count2 == 0:
			c2 = num
			count2++
		default:
			count1--
			count2--
		}
	}

	count1, count2 = 0, 0
	for _, num := range nums {
		if num == c1 {
			count1++
		} else if num == c2 {
			count2++
		}
	}

	res := make([]int, 0, 2)
	maxCount := len(nums) / 3

	if count1 > maxCount {
		res = append(res, c1)
	}
	if count2 > maxCount {
		res = append(res, c2)
	}
	return res
}
