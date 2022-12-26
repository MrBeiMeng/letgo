package code_lists

import (
	"fmt"
	"letgo_repo/utils"
	"strings"
	"syscall"
)

/*
: 负责调用各种方法时提供的基本类型
*/

type Args struct {
	ListNodes []*ListNode
}

func (a Args) IsEmpty() bool {
	if a.ListNodes != nil {
		return false
	}

	return true
}

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

// GetIntListNodeHandler 获取链表
//
//	手动输入
func GetIntListNodeHandler(argNum int) (head *ListNode) {
	numsStr, err := utils.GetInput(fmt.Sprintf("\t参数序号%d\t请输入一个链表,例如:[1,2,3,4]", argNum), 1)
	if err != nil {
		println(err.Error())
		syscall.Exit(-1)
	}

	if strings.Contains(numsStr, "[") {
		numsStr = numsStr[1:]
	}
	if strings.Contains(numsStr, "]") {
		numsStr = numsStr[:len(numsStr)-1]
	}

	numStrArr := strings.Split(numsStr, ",")

	numIntArr := make([]string, 0)
	for _, s := range numStrArr {
		//num, _ := strconv.Atoi(s)
		numIntArr = append(numIntArr, s)
	}

	return GetIntListNode(numIntArr...)
}

// GetIntListNode 获取链表
//
//	通过参数
func GetIntListNode(nums ...string) (head *ListNode) {
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

func GetLinkedList(linkedLists string) (result []*ListNode) {
	if linkedLists == "" {
		return nil
	}

	for _, linkedList := range strings.Split(linkedLists, "],[") {
		linkedList = strings.TrimSpace(linkedList)
		linkedList = strings.ReplaceAll(linkedList, "[", "")
		linkedList = strings.ReplaceAll(linkedList, "]", "")

		// 解析并添加链表
		result = append(result, GetIntListNode(strings.Split(linkedList, ",")...))
	}

	return result
}

// 获取数组

// 获取字符串
