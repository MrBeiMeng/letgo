package service

import (
	"letgo_repo/service/type_def"
)

type CodeServiceI interface {
	Search(wrapper type_def.CodeQueryWrapper) type_def.Questions
	Run(codeNum int, args string)
	InitTodoCode(num int)
}

var CodeService CodeServiceI = CodeServiceImpl{}
