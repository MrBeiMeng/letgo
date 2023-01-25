package todo

import "letgo_repo/system_file/service/todo/type_def"

type ServiceTodo interface {
	CreateSeries(wrapper type_def.SeriesWrapper)
	GetList(wrapper type_def.QueryWrapper) []type_def.TodoSeries
	Save(todo type_def.AddTodo)
}
