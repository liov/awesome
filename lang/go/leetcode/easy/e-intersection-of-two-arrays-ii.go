package leetcode

/*
*
350. 两个数组的交集 II

给定两个数组，编写一个函数来计算它们的交集。

示例 1：

输入：nums1 = [1,2,2,1], nums2 = [2,2]
输出：[2,2]
示例 2:

输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出：[4,9]

说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
*/
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	m1, m2 := make(map[int]int), make(map[int]int)
	for i := range nums1 {
		m1[nums1[i]] = m1[nums1[i]] + 1
	}
	for i := range nums2 {
		m2[nums2[i]] = m2[nums2[i]] + 1
	}
	list := make([]int, 0, len(m1))
	for k, v := range m1 {
		for i := 1; i < min(v, m2[k]); i++ {
			list = append(list, k)
		}
	}
	return list
}
