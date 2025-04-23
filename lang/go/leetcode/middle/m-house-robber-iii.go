package leetcode

/*
*
337. 打家劫舍 III

这个地区只有一个入口，我们称之为“根”。 除了“根”之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果两个直接相连的房子在同一天晚上被打劫，房屋将自动报警。

计算在不触动警报的情况下，小偷一晚能够盗取的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func rob3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var sub1 = root.Val + robNotContain(root.Left) + robNotContain(root.Right)
	var sub2 = rob3(root.Left) + rob3(root.Right)
	return max(sub1, sub2)
}

func robNotContain(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return rob3(root.Left) + rob3(root.Right)
}

func robV2(root *TreeNode) int {
	var rootStatus = dfs(root)
	return max(rootStatus[0], rootStatus[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		return make([]int, 2)
	}
	var l = dfs(node.Left)
	var r = dfs(node.Right)
	var selected = node.Val + l[1] + r[1]
	var notSelected = max(l[0], l[1]) + max(r[0], r[1])
	return []int{selected, notSelected}
}
