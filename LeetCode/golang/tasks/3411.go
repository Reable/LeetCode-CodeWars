package tasks

func MaxLength3411(nums []int) int {

	result := 0

	for l := range nums {
		prod := 1
		gcd := 0
		lcm := 1

		r := l
		for r < len(nums) {
			gcd = gcd3411(gcd, nums[r])
			lcm = (lcm * nums[r]) / gcd3411(lcm, nums[r])
			prod *= nums[r]

			if prod == gcd*lcm && r-l+1 > result {
				result = r - l + 1
			}

			r++
		}

	}

	return result
}

func gcd3411(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
