package todo

import (
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/models"
)

type DATodoImpl struct {
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

func (D *DATodoImpl) SelectAll() (resultList []models.Todo) {
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
