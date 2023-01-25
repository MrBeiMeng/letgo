package question

import (
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/models"
)

type DAQuestionImpl struct {
}

func (D *DAQuestionImpl) GetById(id string) (question models.Question) {
	err := data_access.MysqlDB.Where("frontend_question_id = ?", id).Find(&question).Error
	if err != nil {
		panic(err)
	}

	return
}

func (D *DAQuestionImpl) GetByIds(ids []string) (questions []models.Question) {
	err := data_access.MysqlDB.Where("frontend_question_id in ?", ids).Find(&questions).Error
	if err != nil {
		panic(err)
	}

	return
}
