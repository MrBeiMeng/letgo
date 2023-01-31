package todo

import (
	"letgo_repo/system_file/data_access/manifest"
	manifest_type_def "letgo_repo/system_file/data_access/manifest/type_def"
	"letgo_repo/system_file/data_access/models"
	"letgo_repo/system_file/data_access/question"
	"letgo_repo/system_file/data_access/todo"
	"letgo_repo/system_file/service/todo/type_def"
	"letgo_repo/system_file/utils/enum"
	"letgo_repo/system_file/utils/logger"
	"strings"
)

type ServiceTodoImpl struct{}

func (i ServiceTodoImpl) ChangeDefaultSeries(series string) error {
	return todo.DATodoV1.ChangeDefaultSeries(series)
}

func (i ServiceTodoImpl) GetDefaultSeriesName() (string, error) {
	return todo.DATodoV1.SelectDefaultSeriesName()
}

func (i ServiceTodoImpl) Save(addTodo type_def.AddTodo) {
	//  manifestList 可能是清单 title 名或者 清单 tag 名
	if ok, msg := addTodo.Check(); !ok {
		logger.Logger.Break(msg)
	}

	// 获取系列
	manifests := manifest.DAManifestV1.Select(manifest_type_def.QueryWrapper{
		TitleSlice: addTodo.ManifestList,
		TagSlice:   addTodo.ManifestList,
	})

	// 保存
	todos := make([]models.Todo, 0)
	todoQuestionStatus := make([]models.TodoQuestion, 0)
	for _, manifest := range manifests {
		todos = append(todos, models.Todo{
			ManifestTitle: manifest.Title,
			ManifestMark:  manifest.Mark,
			ManifestTag:   manifest.Tags,
			Series:        addTodo.Series,
		})
	}
	// 保存状态
	todo.DATodoV1.InsertSlice(todos)

	for index, manifest := range manifests {
		for _, frontId := range strings.Split(manifest.QuestionsFrontIds, ",") {
			todoQuestionStatus = append(todoQuestionStatus, models.TodoQuestion{
				TodoId:             todos[index].ID,
				Difficulty:         question.DAQuestionV1.GetById(frontId).Difficulty,
				FrontendQuestionId: frontId,
				Status:             enum.NOT_START,
			})
		}
	}

	todo.DATodoV1.InsertQuestionSlice(todoQuestionStatus)

}

func (i ServiceTodoImpl) GetList(wrapper type_def.QueryWrapper) []type_def.TodoSeries {
	resultList := make([]type_def.TodoSeries, 0)

	todoModels := todo.DATodoV1.SelectAll(models.Todo{Series: wrapper.Series})

	tmpMap := GroupBySeries(todoModels)

	resultList = AppendToResult(tmpMap, resultList)

	return resultList
}

func AppendToResult(tmpMap map[string][]type_def.Todo, resultList []type_def.TodoSeries) []type_def.TodoSeries {
	for key, value := range tmpMap {
		todoSeries := type_def.TodoSeries{
			Series: key,
			Todos:  value,
		}

		resultList = append(resultList, todoSeries)
	}
	return resultList
}

func GroupBySeries(todoModels []models.Todo) map[string][]type_def.Todo {
	tmpMap := make(map[string][]type_def.Todo)
	for _, todoModel := range todoModels {

		// 处理 todoModel
		if _, ok := tmpMap[todoModel.Series]; !ok { // map 中无此key
			tmpMap[todoModel.Series] = make([]type_def.Todo, 0)
		}

		t := type_def.Todo{}
		t.ConvFrom(todoModel)
		todoQuestions := todo.DATodoV1.SelectTodoQuestionsByTodoId(int(todoModel.ID))
		t.TodoQuestions = Conv(todoQuestions, t)

		tmpMap[todoModel.Series] = append(tmpMap[todoModel.Series], t)
	}
	return tmpMap
}

func Conv(todoQuestions []models.TodoQuestion, t type_def.Todo) []type_def.TodoQuestion {
	for _, todoQuestionModel := range todoQuestions {
		var todoQuestion type_def.TodoQuestion
		todoQuestion.FrontendQuestionId = todoQuestionModel.FrontendQuestionId
		todoQuestion.TodoId = todoQuestionModel.TodoId
		todoQuestion.Difficulty = todoQuestionModel.Difficulty
		todoQuestion.Status = todoQuestionModel.Status

		t.TodoQuestions = append(t.TodoQuestions, todoQuestion)
	}

	return t.TodoQuestions
}

func (ServiceTodoImpl) CreateSeries(seriesWrapper type_def.SeriesWrapper) {
	todo.DATodoV1.CreateSeries(seriesWrapper.Series)
}
