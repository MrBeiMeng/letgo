package problems_mapper

import (
	"fmt"
	"gorm.io/gorm/clause"
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/models"
	"letgo_repo/system_file/data_access/todo"
	"letgo_repo/system_file/utils/enum"
	"letgo_repo/system_file/utils/logger"
)

type ProblemsMapperImpl struct {
}

func (p ProblemsMapperImpl) QuestionDone(codeNum string) {
	updateSql := "update questions set status = 'AC' where frontend_question_id = ?;"

	err := data_access.MysqlDB.Exec(updateSql, codeNum).Error
	if err != nil {
		panic(err)
	}

	updateSql2 := "update todo_questions,todos set todo_questions.status = ? where todo_questions.todo_id = todos.id and todos.series = ? and todo_questions.frontend_question_id = ?"

	series, err := todo.DATodoV1.SelectDefaultSeriesName()
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("使用默认series [%s]", series)

	err = data_access.MysqlDB.Exec(updateSql2, enum.DONE, series, codeNum).Error
	if err != nil {
		panic(err)
	}
}

func (p ProblemsMapperImpl) GetTests(codeNum string) (result []models.QuestionTest) {
	err := data_access.MysqlDB.Where("frontend_question_id = ?", codeNum).Find(&result).Error
	if err != nil {
		panic(err)
	}

	return
}

func (p ProblemsMapperImpl) SaveOrUpdateTest(codeNum int, strArgs string, rightAnswer string) error {
	var test models.QuestionTest
	test.FrontendQuestionId = fmt.Sprintf("%d", codeNum)
	test.Args = strArgs
	test.RightAnswer = rightAnswer

	err := data_access.MysqlDB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&test).Error
	return err
}

func (p ProblemsMapperImpl) CountDone(codeNums []string) (num int) {
	selectSql := "select count(1) from questions where status = 'AC' and frontend_question_id in ? ;"

	err := data_access.MysqlDB.Raw(selectSql, codeNums).Scan(&num).Error
	if err != nil {
		panic(err)
	}

	return num
}

func (p ProblemsMapperImpl) GetTodos() (resultList []models.ToDoQuestion) {
	err := data_access.MysqlDB.Order("sort").Find(&resultList).Error
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

	err := data_access.MysqlDB.Save(&operate).Error
	if err != nil {
		panic(err)
	}
}

func (p ProblemsMapperImpl) InitInsertQuestionStatus(num int) {
	insertSql := "-"
	err := data_access.MysqlDB.Exec(insertSql, num, enum.TODO).Error
	if err != nil {
		panic(err)
	}
}

func (p ProblemsMapperImpl) GetByCodeNumInDB(codeNum string) (result models.Question) {
	var question models.Question
	err := data_access.MysqlDB.Where("frontend_question_id = ?", codeNum).First(&question).Error // 根据整型主键查找
	if err != nil {
		panic(err)
	}

	return question
}

func (p ProblemsMapperImpl) GetByCodeNum(codeNum int) (question models.Question) {
	err := data_access.MysqlDB.Where("frontend_question_id = ?", codeNum).Find(&question).Error
	if err != nil {
		panic(err)
	}

	err = data_access.MysqlDB.Model(&question).Association("TopicTags").Find(&question.TopicTags)
	if err != nil {
		panic(err)
	}

	err = data_access.MysqlDB.Model(&question).Association("TopCompanyTags").Find(&question.TopCompanyTags)
	if err != nil {
		panic(err)
	}

	return question
}
