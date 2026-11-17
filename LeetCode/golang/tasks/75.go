package tasks

func SortColors75(nums []int) {
	if len(nums) < 2 {
		return
	}

	pivot := nums[len(nums)/2]

	l := 0
	r := len(nums) - 1
	for l <= r {
		for nums[l] < pivot {
			l++
		}
		for nums[r] > pivot {
			r--
		}

		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	if r > 0 {
		SortColors75(nums[:r+1])
	}

	if l < len(nums) {
		SortColors75(nums[l:])
	}
}
