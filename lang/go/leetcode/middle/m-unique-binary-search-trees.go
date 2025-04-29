package leetcode

/**
不同的二叉搜索树

给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

1         3     3      2      1
\       /     /      / \      \
3     2     1      1   3      2
/     /       \                 \
2     1         2                 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/unique-binary-search-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//Gn = sum(i in 1..n){G(i-1)G(n-i)}
//Gn+1=(2(2n+1)/(n+2))Gn
func numTrees(n int) int {
	var c = 1
	for i := range n {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}

func numTreesV2(n int) int {
	var arr = make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			arr[i] += arr[j-1] * arr[i-j]
		}
	}
	return arr[n]
}

/**
 * 生成二叉搜索树
 */
func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: n}}
	}
	return helper(1, n)
}

func helper(start int, end int) []*TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return []*TreeNode{&TreeNode{Val: start}}
	}
	var list = make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		var left = helper(start, i-1)
		var right = helper(i+1, end)
		for l := range left {
			for r := range right {
				var tree = &TreeNode{Val: i}
				tree.Left = left[l]
				tree.Right = right[r]
				list = append(list, tree)
			}
		}
	}
	return list
}
