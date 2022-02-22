package single

func SingleNumber(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length%2 == 0 {
		return -1
	}
	ans := 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
