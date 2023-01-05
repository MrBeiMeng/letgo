package data_access

import (
	"letgo_repo/letgo_file/data_access/models"
)

type ProblemsMapperI interface {
	GetByCodeNum(codeNum int) (question models.Questions)
	GetByCodeNumInDB(codeNum int) models.QuestionInfo
	InitInsertQuestionStatus(num int)
	OperationLog(summary, msg, opType string)
	GetTodos() []models.ToDoQuestion
	CountDone(codeNums []string) int
	SaveAnswer(codeNum int, strArgs string, rightAnswer string) error
}

var ProblemsMapper ProblemsMapperI = ProblemsMapperImpl{}
