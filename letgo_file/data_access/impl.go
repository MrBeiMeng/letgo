package data_access

import (
	"letgo_repo/letgo_file/data_access/models"
	"letgo_repo/letgo_file/utils/enum"
)

type ProblemsMapperImpl struct {
}

func (p ProblemsMapperImpl) CountDone(codeNums []string) (num int) {
	selectSql := "select count(1) from questions where status = 'AC' and frontend_question_id in ? ;"

	err := MysqlDB.Raw(selectSql, codeNums).Scan(&num).Error
	if err != nil {
		panic(err)
	}

	return num
}

func (p ProblemsMapperImpl) GetTodos() (resultList []models.ToDoQuestion) {
	err := MysqlDB.Order("sort").Find(&resultList).Error
	if err != nil {
		panic(err)
	}

	return
}

func (p ProblemsMapperImpl) OperationLog(summary, msg, opType string) {
	var operate models.OperationRecords

	operate.Summary = summary
	operate.Msg = msg
	operate.OpType = opType

	err := MysqlDB.Save(&operate).Error
	if err != nil {
		panic(err)
	}
}

func (p ProblemsMapperImpl) InitInsertQuestionStatus(num int) {
	insertSql := "-"
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
