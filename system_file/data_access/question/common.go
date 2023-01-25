package question

import "letgo_repo/system_file/data_access/models"

type DAQuestion interface {
	GetByIds([]string) []models.Question
	GetById(string) models.Question
}

var DAQuestionV1 DAQuestion = &DAQuestionImpl{}
