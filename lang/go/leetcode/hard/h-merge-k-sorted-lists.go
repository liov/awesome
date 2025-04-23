package leetcode

import (
	"iter"
	"test/leetcode"
	easy "test/leetcode/easy"
)

/**
23. 合并K个升序链表

给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。



示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
1->4->5,
1->3->4,
2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]


提示：

k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//if 要加括号就很不爽
func mergeKLists(lists []ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return &lists[0]
	}
	tree := NewBinaryTree()
	var tmp *ListNode
	for i := range lists {
		tmp = &lists[i]
		for tmp != nil {
			tree.insert(tmp.Val)
			tmp = tmp.Next
		}
	}
	var ans = &ListNode{}
	tmp = ans
	for v := range tree.Sequence() {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return ans.Next
}

type BinaryTree struct {
	root *TreeNode
	size int
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (t *BinaryTree) isEmpty() bool {
	return t.size == 0
}

func (t *BinaryTree) insert(v int) {
	t.size += 1
	if t.root == nil {
		t.root = &TreeNode{Val: v}
		return
	}
	t.root.Insert(v)
}

func (t *BinaryTree) insertLoop(v int) {
	t.size += 1
	if t.root == nil {
		t.root = &TreeNode{Val: v}
		return
	}
	var tmp = t.root
	for true {
		if v > tmp.Val {
			if tmp.Right == nil {
				tmp.Right = &TreeNode{Val: v}
				return
			} else {
				tmp = tmp.Right
			}
		} else {
			if tmp.Left == nil {
				tmp.Left = &TreeNode{Val: v}
				return
			} else {
				tmp = tmp.Left
			}
		}
	}
}

/**
 * 前序遍历
 * @param node 开始遍历元素
 */
func (t *BinaryTree) prevRecursive(node *TreeNode) {
	if node != nil {
		print("${node.value} ")
		t.prevRecursive(node.Left)
		t.prevRecursive(node.Right)
	}
}

func (t *BinaryTree) prevIterator(node *TreeNode) []int {
	var result []int
	stack := leetcode.NewStack[*TreeNode]()
	if t.root != nil {
		stack.Push(t.root)
	} else {
		return result
	}

	var tmp *TreeNode
	for len(stack) > 0 {
		tmp, _ = stack.Pop()
		result = append(result, tmp.Val)
		if tmp.Right != nil {
			stack.Push(tmp.Right) // 右节点入栈
		}
		if tmp.Left != nil {
			stack.Push(tmp.Left) // 左节点入栈
		}
	}
	return result
}

/**
 * 中序遍历
 * @param node 开始遍历元素
 */
func (t *BinaryTree) midRecursive(node *TreeNode) {
	if node != nil {
		t.midRecursive(node.Left)
		print("${node.value} ")
		t.midRecursive(node.Right)
	}
}

func (t *BinaryTree) midIterator() []int {
	var result []int
	stack := leetcode.NewStack[*TreeNode]()
	var tmp = t.root
	for len(stack) > 0 || tmp != nil {
		for tmp != nil {
			stack.Push(tmp) // 添加根节点
			tmp = tmp.Left  // 循环添加左节点
		}
		tmp, _ = stack.Pop()
		result = append(result, tmp.Val)
		tmp = tmp.Right
	}
	return result
}

func (t *BinaryTree) Sequence() iter.Seq[int] {
	return func(yield func(int) bool) {
		var stack = leetcode.NewStack[*TreeNode]()
		var tmp = t.root
		for len(stack) > 0 || tmp != nil {
			for tmp != nil {
				stack.Push(tmp) // 添加根节点
				tmp = tmp.Left  // 循环添加左节点
			}
			tmp, _ = stack.Pop()
			yield(tmp.Val)
			tmp = tmp.Right
		}
	}
}

/**
 * 后序遍历
 * @param node 开始遍历元素
 */
func subRecursive(node *TreeNode) {
	if node != nil {
		subRecursive(node.Left)
		subRecursive(node.Right)
		print("${node.value} ")
	}
}

func mergeKListsV2(lists []ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []ListNode, l int, r int) *ListNode {
	if l == r {
		return &lists[l]
	}
	if l > r {
		return nil
	}
	var mid = (l + r) >> 1
	return easy.MergeTwoListsV2(merge(lists, l, mid), merge(lists, mid+1, r))
}
