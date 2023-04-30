package letgo_2_add_two_numbers

import (
	. "letgo_repo/system_file/code_enter"
)

/*两数相加 | https://leetcode.cn/problems/add-two-numbers*/

/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//type ListNode IntListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	upNum := 0 // 进一位置

	head := &ListNode{}
	point := head

	// 从前往后去加，逢十进一
	for true {
		if l1 == nil && l2 == nil && upNum == 0 { // 当两个链表都跑完时退出
			break
		}
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		point.Next = &ListNode{
			Val:  (num1 + num2 + upNum) % 10,
			Next: nil,
		}
		point = point.Next

		upNum = (num1 + num2 + upNum) / 10
	}

	return head.Next
}
