package letgo_merge_two_sorted_lists

import (
	. "letgo_repo/system_file/code_enter"
	"math"
)

/*合并两个有序链表 | https://leetcode.cn/problems/merge-two-sorted-lists*/

/*
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例 1：
输入：l1 = [1,2,4], l2 = [1,3,4]
输出：[1,1,2,3,4,4]
示例 2：

输入：l1 = [], l2 = []
输出：[]
示例 3：

输入：l1 = [], l2 = [0]
输出：[0]

提示：
两个链表的节点数目范围是 [0, 50]
-100 <= Node.val <= 100
l1 和 l2 均按 非递减顺序 排列
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode = &ListNode{}
	tmp := head

	for !bothNil(list1, list2) {
		val1 := getLastPointOrMaxInt(list1)
		val2 := getLastPointOrMaxInt(list2)

		if val1 <= val2 {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}

		tmp = tmp.Next
	}

	return head.Next
}

func getLastPointOrMaxInt(list1 *ListNode) int {
	val1 := math.MaxInt
	if list1 != nil {
		val1 = list1.Val
	}
	return val1
}

func bothNil(list1 *ListNode, list2 *ListNode) bool {
	return list1 == nil && list2 == nil
}
