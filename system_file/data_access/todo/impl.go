package todo

import (
	"errors"
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/models"
)

type DATodoImpl struct {
}

func (D *DATodoImpl) ChangeDefaultSeries(series string) error {

	updateSql := "update todos set default = null where dafalue = true;"

	err := data_access.MysqlDB.Exec(updateSql).Error
	if err != nil {
		return err
	}

	updateSql = "update todos set default = true where series = ?;"
	err = data_access.MysqlDB.Exec(updateSql, series).Error
	if err != nil {
		return err
	}

	return nil
}

func (D *DATodoImpl) SelectDefaultSeriesName() (string, error) {

	selectSql := "select * from todos where `default` is true;"

	resultList := make([]models.Todo, 0)

	err := data_access.MysqlDB.Raw(selectSql).Find(&resultList).Error
	if err != nil {
		panic(err)
	}

	if len(resultList) == 0 {
		return "", errors.New("没有default")
	}

	return resultList[0].Series, nil
}

func (D *DATodoImpl) InsertQuestionSlice(todoQuestions []models.TodoQuestion) {
	err := data_access.MysqlDB.Create(&todoQuestions).Error
	if err != nil {
		panic(err)
	}
}

func (D *DATodoImpl) InsertSlice(todos []models.Todo) {
	err := data_access.MysqlDB.Create(&todos).Error
	if err != nil {
		panic(err)
	}
}

func (D *DATodoImpl) SelectTodoQuestionsByTodoId(todoId int) (resultList []models.TodoQuestion) {
	err := data_access.MysqlDB.Where("todo_id = ?", todoId).Find(&resultList).Error
	if err != nil {
		panic(err)
	}

	return
}

func (D *DATodoImpl) SelectAll(queryBody models.Todo) (resultList []models.Todo) {
	err := data_access.MysqlDB.Find(&resultList).Error
	if err != nil {
		panic(err)
	}

	return
}

func (D *DATodoImpl) CreateSeries(series string) {
	modelTodo := models.Todo{Series: series}

	err := data_access.MysqlDB.Create(&modelTodo).Error
	if err != nil {
		panic(err)
	}
}
