package service

import (
	"letgo_repo/service/type_def"
)

type CodeServiceI interface {
	Search(wrapper type_def.CodeQueryWrapper) type_def.Questions
	Run(codeNum int, args string)
	InitTodoCode(num int)
	GetByCodeNum(num int) type_def.Question
}

var CodeService CodeServiceI = CodeServiceImpl{}
