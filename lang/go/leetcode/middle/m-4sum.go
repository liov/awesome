package leetcode

import "sort"

/**
18. 四数之和

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
[-1,  0, 0, 1],
[-2, -1, 1, 2],
[-2,  0, 0, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 哈哈哈哈，执行用时 : 620 ms , 在所有 Kotlin 提交中击败了 8.00% 的用户
func fourSum(nums []int, target int) [][]int {
	return nSum(nums, target, 4)
}

func nSum(nums []int, target int, n int) [][]int {
	sort.Ints(nums)
	if len(nums) == 0 || (target > 0 && nums[0] > target || target < 0 && nums[len(nums)-1] < target) {
		return [][]int{}
	}
	var m = make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	var ans = make([][]int, 0)
	var tmp = make([]int, n)
	nSumHelper(nums, 0, target, n, m, &ans, tmp)
	return ans
}

// subList改为起始结束位置
func nSumHelper(nums []int, subStart int, target int, n int, m map[int]int,
	ans *[][]int, tmp []int) {
	if n == 1 {
		if v, ok := m[target]; ok && v >= subStart {
			tmp[len(tmp)-n] = target
			*ans = append(*ans, tmp)
		}
		return
	}
	for i := subStart; i < len(nums); i++ {
		if i > subStart && nums[i-1] == nums[i] {
			continue
		}
		if target > 0 && nums[i] > target || target < 0 && nums[len(nums)-1] < target {
			break
		}
		tmp[len(tmp)-n] = nums[i]
		nSumHelper(nums, i+1, target-nums[i], n-1, m, ans, tmp)
	}
}
