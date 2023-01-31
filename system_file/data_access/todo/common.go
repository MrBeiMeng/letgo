package todo

import "letgo_repo/system_file/data_access/models"

type DATodo interface {
	CreateSeries(series string)
	InsertSlice(todos []models.Todo)
	InsertQuestionSlice(todos []models.TodoQuestion)
	SelectAll(models.Todo) []models.Todo
	SelectTodoQuestionsByTodoId(todoIds int) []models.TodoQuestion
	SelectDefaultSeriesName() (string, error)
	ChangeDefaultSeries(series string) error
}

var DATodoV1 DATodo = &DATodoImpl{}
