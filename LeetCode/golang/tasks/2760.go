package tasks

func LongestAlternatingSubarray2760(nums []int, threshold int) int {

	count := 0

	for l := 0; l < len(nums); l++ {
		if nums[l]%2 != 0 || nums[l] > threshold {
			continue
		}
		tempCount := 1

		for r := l + 1; r < len(nums); r++ {
			if nums[r] > threshold {
				break
			}

			if nums[r-1]%2 == nums[r]%2 {
				break
			}
			tempCount++
			l = r - 1
		}
		if tempCount > count {
			count = tempCount
		}
	}

	return count
}

func LongestAlternatingSubarray2760_v2(nums []int, threshold int) int {
	count := 0
	maxCount := 0

	for i := range nums {
		if nums[i] > threshold {
			count = 0
			continue
		}

		if count == 0 {
			if nums[i]%2 == 0 {
				count = 1
			}
		} else if nums[i-1]%2 != nums[i]%2 {
			count++
		} else {
			if nums[i]%2 == 0 {
				count = 1
			} else {
				count = 0
			}
		}

		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
