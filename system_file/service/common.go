package service

import (
	"letgo_repo/system_file/service/type_def"
)

type CodeServiceI interface {
	Search(wrapper type_def.CodeQueryWrapper) type_def.Questions
	Run(arg type_def.RunWrapper)
	InitTodoCode(num int)
	GetByCodeNum(num int) type_def.Question
	OperateLog(summary, msg, opType string)
	GetToDos() []type_def.ToDoQuestion
}

var CodeService CodeServiceI = CodeServiceImpl{}
