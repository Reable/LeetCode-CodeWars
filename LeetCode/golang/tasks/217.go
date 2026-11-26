package tasks

func ContainsDuplicate217(nums []int) bool {
	temp := make(map[int]int)

	for _, v := range nums {
		if _, ok := temp[v]; ok {
			return true
		} else {
			temp[v]++
		}
	}

	return false
}
