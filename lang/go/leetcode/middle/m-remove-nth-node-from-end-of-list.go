package leetcode

/*
*
19. 删除链表的倒数第 N 个结点

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}
	var nNode = head
	var tmp = head
	var count = 0
	for tmp.Next != nil {
		tmp = tmp.Next
		count++
		if count > n {
			nNode = nNode.Next
		}
	}
	if count+1 == n {
		return head.Next
	}
	nNode.Next = nNode.Next.Next
	return head
}
