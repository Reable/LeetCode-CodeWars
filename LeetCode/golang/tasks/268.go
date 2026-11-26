package tasks

func MissingNumber268(nums []int) int {
	maxV := 0
	temp := make([]int, len(nums)+1)
	for i := range temp {
		temp[i] = -1
	}

	for _, n := range nums {
		temp[n] = n
		if n > maxV {
			maxV = n
		}
	}

	for i := range temp {
		if temp[i] == -1 {
			return i
		}
	}
	return -1
}
