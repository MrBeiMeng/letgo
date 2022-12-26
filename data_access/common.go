package data_access

import (
	"letgo_repo/data_access/models"
)

type ProblemsMapperI interface {
	GetByCodeNum(codeNum int) (question models.Question, questionStatus models.QuestionStatus)
	GetByCodeNumInDB(CodeNum int) models.Question
}

var ProblemsMapper ProblemsMapperI = ProblemsMapperImpl{}
