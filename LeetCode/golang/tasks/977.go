package tasks

func SortedSquares977(nums []int) []int {
	result := make([]int, len(nums))
	l := 0
	r := len(nums) - 1
	i := r

	for l <= r {
		left := nums[l] * nums[l]
		right := nums[r] * nums[r]

		if left > right {
			result[i] = left
			l++
		} else {
			result[i] = right
			r--
		}
		i--
	}
	return result
}
