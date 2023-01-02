package data_access

import (
	"letgo_repo/data_access/models"
)

type ProblemsMapperI interface {
	GetByCodeNum(codeNum int) (question models.Questions)
	GetByCodeNumInDB(codeNum int) models.QuestionInfo
	InitInsertQuestionStatus(num int)
	OperationLog(summary, msg, opType string)
	GetTodos() []models.ToDoQuestion
	CountDone(codeNums []string) int
}

var ProblemsMapper ProblemsMapperI = ProblemsMapperImpl{}
