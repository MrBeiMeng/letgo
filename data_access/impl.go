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

func (p ProblemsMapperImpl) GetByCodeNum(codeNum int) (question models.Question) {
	err := MysqlDB.Raw("select questions.id, title, title_slug, article_live, article_slug, level, total_submitted, total_acs, frontend_id, translated_title, content, translated_content, code_snippets, question_id, star, status, visible from questions,question_status where questions.id = question_status.question_id and question_id = ?", codeNum).Scan(&question).Error
	if err != nil {
		panic(err)
	}

	return question
}
