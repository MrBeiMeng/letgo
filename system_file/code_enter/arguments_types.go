package code_enter

import (
	"fmt"
	"strconv"
	"strings"
)

/*
: 负责调用各种方法时提供的基本类型
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	println(l.Sprint())
}

func (l *ListNode) Sprint() string {
	tmpNode := l
	answerNums := make([]string, 0)
	for tmpNode != nil {
		answerNums = append(answerNums, strconv.Itoa(tmpNode.Val))
		tmpNode = tmpNode.Next
	}

	pStr := strings.Join(answerNums, ",")
	return fmt.Sprintf("[%s]\t", pStr)
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
		num, _ := strconv.Atoi(nums[i])
		tmpHead.Val = num
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
	//if linkedList == "" {
	//	return nil
	//}

	linkedList = strings.Trim(linkedList, "[] ")
	if linkedList == "" {
		return nil
	}
	result = a.GetIntListNode(strings.Split(linkedList, ",")...)

	return result
}

// GetStringArr 获取数组
func (a ArgsHandler) GetStringArr(s string) []string {
	s = strings.Trim(s, "][")

	strList := make([]string, 0)
	for _, str := range strings.Split(s, ",") {
		if strings.EqualFold(str, "") {
			continue
		}

		strList = append(strList, str)
	}

	return strList
}

// GetIntArr 获取数组
func (a ArgsHandler) GetIntArr(s string) []int {
	s = strings.Trim(s, "][")

	nums := make([]int, 0)
	for _, str := range strings.Split(s, ",") {
		if strings.EqualFold(str, "") {
			continue
		}

		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	return nums
}

// GetIntMatrix 获取2x2矩阵
func (a ArgsHandler) GetIntMatrix(s string) [][]int {
	matrix := make([][]int, 0)

	if strings.EqualFold(s, "[]") {
		return matrix
	}

	for _, arrStr := range strings.Split(s, "],[") {
		matrix = append(matrix, a.GetIntArr(arrStr))
	}

	return matrix
}

func (a ArgsHandler) GetInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}
