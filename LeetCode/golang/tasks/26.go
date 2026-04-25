package tasks

func RemoveDuplicates26(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 1
	fast := 1

	for fast < len(nums) {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	return slow
}
