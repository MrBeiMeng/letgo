package code_lists

import (
	"fmt"
	"strconv"
	"strings"
)

/*
: 负责调用各种方法时提供的基本类型
*/

type ListNode struct {
	Val  string
	Next *ListNode
}

func (l *ListNode) Print() {
	tmpNode := l
	answerNums := make([]string, 0)
	for tmpNode != nil {
		answerNums = append(answerNums, tmpNode.Val)
		tmpNode = tmpNode.Next
	}

	pStr := strings.Join(answerNums, ",")
	fmt.Printf("[%s]\t", pStr)
}

// 获取字符串

var ArgsHandlerV1 ArgsHandler = ArgsHandler{}

type ArgsHandler struct {
}

// GetIntListNode 获取链表
//
//	通过参数
func (a ArgsHandler) GetIntListNode(nums ...string) (head *ListNode) {
	var tmpHead *ListNode
	tmpHead = &ListNode{}
	head = tmpHead

	for i := 0; i < len(nums); i++ {
		tmpHead.Val = nums[i]
		if i+1 < len(nums) {
			nextHead := ListNode{}
			tmpHead.Next = &nextHead
			tmpHead = tmpHead.Next
		}
	}

	return head
}

func (a ArgsHandler) GetLinkedLists(linkedLists string) (result []*ListNode) {
	if linkedLists == "" {
		return nil
	}

	for _, linkedList := range strings.Split(linkedLists, "],[") {
		linkedList = strings.Trim(linkedList, "[] ")

		// 解析并添加链表
		result = append(result, a.GetIntListNode(strings.Split(linkedList, ",")...))
	}

	return result
}

func (a ArgsHandler) GetLinkedList(linkedList string) (result *ListNode) {
	if linkedList == "" {
		return nil
	}

	linkedList = strings.Trim(linkedList, "[] ")
	result = a.GetIntListNode(strings.Split(linkedList, ",")...)

	return result
}

// GetIntArr 获取数组
func (a ArgsHandler) GetIntArr(s string) []int {
	s = strings.Trim(s, "][")

	nums := make([]int, 0)
	for _, str := range strings.Split(s, ",") {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	return nums
}

func (a ArgsHandler) GetInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}
