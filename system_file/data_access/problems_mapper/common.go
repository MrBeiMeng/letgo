package problems_mapper

import (
	"letgo_repo/system_file/data_access/models"
)

type ProblemsMapperI interface {
	GetByCodeNum(codeNum int) (question models.Question)
	GetByCodeNumInDB(codeNum string) models.Question
	InitInsertQuestionStatus(num int)
	OperationLog(summary, msg, opType string)
	GetTodos() []models.ToDoQuestion
	CountDone(codeNums []string) int
	SaveOrUpdateTest(codeNum int, strArgs string, rightAnswer string) error
	GetTests(codeNum string) []models.QuestionTest
	QuestionDone(codeNum string)
}

var ProblemsMapper ProblemsMapperI = ProblemsMapperImpl{}
