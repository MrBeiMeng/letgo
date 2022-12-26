package service

import (
	"letgo_repo/code_lists"
	"letgo_repo/service/type_def"
)

type CodeServiceI interface {
	Search(wrapper type_def.CodeQueryWrapper) code_lists.CodeChallengeListObj
	GetLinkedList(linkedLists string) (result []*code_lists.ListNode)
	Run(codeNum int, args type_def.Args)
	RunDemo(codeNum int)
	SearchInDBByNo(codeNum int) code_lists.CodeInfo
}

var CodeService CodeServiceI = CodeServiceImpl{}
