package tasks

type ProductOfNumbers1352 struct {
	nums []int
}

func Constructor1352() ProductOfNumbers1352 {
	return ProductOfNumbers1352{
		nums: []int{},
	}
}

func (this *ProductOfNumbers1352) Add1352(num int) {
	this.nums = append(this.nums, num)
}

func (this *ProductOfNumbers1352) GetProduct1352(k int) int {
	sum := 1
	for i := len(this.nums) - 1; i >= len(this.nums)-k; i-- {
		sum *= this.nums[i]
	}
	return sum
}
