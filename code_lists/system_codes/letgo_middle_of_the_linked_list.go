package system_codes

import (
	. "letgo_repo/system_file/code_enter"
)

/*链表的中间结点 | https://leetcode.cn/problems/middle-of-the-linked-list*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	fp := head
	sp := head

	for fp != nil && fp.Next != nil {
		sp = sp.Next
		fp = fp.Next.Next
	}

	return sp
}
