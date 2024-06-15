package leetcode

// 只出现一次的数字 https://leetcode-cn.com/problems/single-number/
func singleNumber(nums []int) int {
	x := nums[0]
	for i := 1; i < len(nums); i++ {
		x = x ^ nums[i]
	}
	return x
}
