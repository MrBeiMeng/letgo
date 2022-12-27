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

func (p ProblemsMapperImpl) GetByCodeNumInDB(codeNum int) (result models.Question) {
	var question models.Question
	MysqlDB.First(&question, codeNum) // 根据整型主键查找

	return question
}

func (p ProblemsMapperImpl) GetByCodeNum(codeNum int) (question models.Question, questionStatus models.QuestionStatus) {
	MysqlDB.First(&question, codeNum) // 根据整型主键查找

	err := MysqlDB.Where("question_id=?", codeNum).Table("question_status").Find(&questionStatus).Error
	if err != nil {
		panic(err)
	}

	return question, questionStatus
}

//func convToCodeInfo(question models.Question, questionStatus models.QuestionStatus) code_lists.CodeInfo {
//  result.Title = q.TranslatedTitle
//	result.CodeNum, _ = strconv.Atoi(q.Id)
//	result.Level = q.Level
//	result.Description = q.TranslatedContent
//	result.Visible = true
//	result.Url = "https://leetcode.cn/problems/" + q.TitleSlug
//	codeInfo := question.ConvQuestionToCodeInfo()
//	codeInfo.Star = questionStatus.Star
//	codeInfo.Status = questionStatus.Status
//	codeInfo.Visible = questionStatus.Visible
//	return codeInfo
//}
