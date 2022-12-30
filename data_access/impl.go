package data_access

import (
	"letgo_repo/data_access/models"
	"letgo_repo/utils/enum"
)

type ProblemsMapperImpl struct {
}

func (p ProblemsMapperImpl) InitInsertQuestionStatus(num int) {
	insertSql := "insert into letgo.question_status (question_id, status) values (?,?);"
	err := MysqlDB.Exec(insertSql, num, enum.TODO).Error
	if err != nil {
		panic(err)
	}
}

func (p ProblemsMapperImpl) GetByCodeNumInDB(codeNum int) (result models.QuestionInfo) {
	var question models.QuestionInfo
	MysqlDB.First(&question, codeNum) // 根据整型主键查找

	return question
}

func (p ProblemsMapperImpl) GetByCodeNum(codeNum int) (question models.Questions) {
	err := MysqlDB.Where("frontend_question_id = ?", codeNum).First(&question).Error
	if err != nil {
		panic(err)
	}

	err = MysqlDB.Model(&question).Association("TopicTags").Find(&question.TopicTags)
	if err != nil {
		panic(err)
	}

	err = MysqlDB.Model(&question).Association("TopCompanyTags").Find(&question.TopCompanyTags)
	if err != nil {
		panic(err)
	}

	return question
}
