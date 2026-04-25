package tasks

func RemoveElement27(nums []int, val int) int {
	if len(nums) < 1 {
		return 0
	}

	slow := 0
	fast := 0

	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
