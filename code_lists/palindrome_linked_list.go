package code_lists

func isPalindrome(head *ListNode) bool {
	length := getLen(head)
	if length == 0 {
		return false
	}

	tmpHead := head
	var reversedHead *ListNode
	midIndex := length / 2

	for i := 0; i < midIndex; i++ {
		reversedHead, tmpHead = reverseIn(reversedHead, tmpHead)
	}

	if length%2 != 0 {
		tmpHead = tmpHead.Next
	}

	return isSameLink(tmpHead, reversedHead)
}

func isSameLink(link1 *ListNode, link2 *ListNode) bool {
	for link1 != nil && link2 != nil {
		if link1.Val != link2.Val {
			return false
		}
		link1 = link1.Next
		link2 = link2.Next
	}
	if link1 == nil && link2 == nil {
		return true
	}

	return false
}

func reverseIn(reHead *ListNode, head *ListNode) (*ListNode, *ListNode) {
	if reHead == nil {
		reHead = head
		head = head.Next
		reHead.Next = nil
	} else {
		tmpNode := head
		head = head.Next
		tmpNode.Next = reHead
		reHead = tmpNode
	}

	return reHead, head
}

func getLen(head *ListNode) (length int) {
	tmpHead := head
	for tmpHead != nil {
		length++
		tmpHead = tmpHead.Next
	}
	return length
}

type PalindromeLinkedList struct {
}

func (p PalindromeLinkedList) GetTest() string {
	return "[1,2,3,2,1]"
}

func (p PalindromeLinkedList) GetFunc() interface{} {
	return isPalindrome
}

func (p PalindromeLinkedList) GetCodeNum() int {
	return 234
}

//
//func (p PalindromeLinkedList) GetTags() []string {
//	return []string{enum.LINKED_LIST}
//}
//
//func (p PalindromeLinkedList) RunDemo() {
//	print("\t参数")
//	//head := GetIntListNode([]string{"1","2","3"}...)
//	heads := ArgsHandlerV1.GetLinkedList("[1,2,3,2,1]")
//	heads[0].Print()
//
//	print("\t结果")
//	println(isPalindrome(heads[0]))
//}
//
//func (p PalindromeLinkedList) Run(args Args) {
//	//head := GetIntListNodeHandler(1)
//	print("\t参数")
//	heads := args.ListNodes
//	heads[0].Print()
//
//	print("\t结果")
//	println(isPalindrome(heads[0]))
//}
