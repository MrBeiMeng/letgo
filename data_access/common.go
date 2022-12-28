package data_access

import (
	"letgo_repo/data_access/models"
)

type ProblemsMapperI interface {
	GetByCodeNum(codeNum int) (question models.Question)
	GetByCodeNumInDB(CodeNum int) models.QuestionInfo
	InitInsertQuestionStatus(num int)
}

var ProblemsMapper ProblemsMapperI = ProblemsMapperImpl{}
