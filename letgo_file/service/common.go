package service

import (
	"letgo_repo/letgo_file/service/type_def"
)

type CodeServiceI interface {
	Search(wrapper type_def.CodeQueryWrapper) type_def.Questions
	Run(codeNum int, args string)
	InitTodoCode(num int)
	GetByCodeNum(num int) type_def.Question
	OperateLog(summary, msg, opType string)
	GetToDos() []type_def.ToDoQuestion
}

var CodeService CodeServiceI = CodeServiceImpl{}
