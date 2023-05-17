package letgo_merge_k_sorted_lists

import (
	. "letgo_repo/system_file/code_enter"
	"math"
)

/*合并 K 个升序链表 | https://leetcode.cn/problems/merge-k-sorted-lists*/

/*
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
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	answerHead := &ListNode{}
	tmp := answerHead

	var minIndex int
	for minIndex != -1 {
		// 思路，每次寻找一个最小的节点，用lists的索引表示
		minIndex = getMinIndex(lists)
		if minIndex == -1 {
			break
		}

		// 将这个节点安装到结果链表中，让那个链表头向后移动一位，断掉它的尾巴。
		tmp = addToAnswer(lists, tmp, minIndex)
	}

	return answerHead.Next
}

func getMinIndex(lists []*ListNode) int {
	minValue := math.MaxInt
	answerIndex := -1

	for i, link := range lists {
		if link == nil {
			continue
		}

		if link.Val < minValue {
			minValue = link.Val
			answerIndex = i
		}

	}

	return answerIndex
}

func addToAnswer(lists []*ListNode, tmp *ListNode, minIndex int) *ListNode {
	tmp.Next = lists[minIndex]
	lists[minIndex] = lists[minIndex].Next
	tmp.Next.Next = nil
	tmp = tmp.Next

	return tmp
}
