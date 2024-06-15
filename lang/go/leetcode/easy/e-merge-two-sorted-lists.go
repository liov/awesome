package leetcode

/**
 *
21. 合并两个有序链表

将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

https://leetcode-cn.com/problems/merge-two-sorted-lists/
*/
//太简陋了这个连表，本来还想实现Iterator

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node1 = l1
	var node2 = l2
	var headNode *ListNode
	if l1.Val < l2.Val {
		headNode = &ListNode{Val: l1.Val}
		node1 = l1.Next
	} else {
		headNode = &ListNode{Val: l2.Val}
		node2 = l2.Next
	}
	ans := headNode
	for {
		if node1 == nil {
			headNode.Next = node2
			return ans
		}
		if node2 == nil {
			headNode.Next = node1
			return ans
		}
		if node1.Val < node2.Val {
			headNode.Next = &ListNode{Val: node1.Val}
			headNode = headNode.Next
			node1 = node1.Next
		} else {
			headNode.Next = &ListNode{Val: node2.Val}
			headNode = headNode.Next
			node2 = node2.Next
		}
	}
}

func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var node1 = l1
	var node2 = l2
	var headNode = &ListNode{}
	ans := headNode
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			headNode.Next = node1
			node1 = node1.Next
		} else {
			headNode.Next = node2
			node2 = node2.Next
		}
		headNode = headNode.Next
	}
	if node1 == nil {
		headNode.Next = node2
	} else {
		headNode.Next = node1
	}
	return ans.Next
}
