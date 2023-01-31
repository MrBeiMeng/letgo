package type_def

import (
	"letgo_repo/system_file/data_access/models"
	"strings"
)

type SeriesWrapper struct {
	Series string
}

type QueryWrapper struct {
	Series string
}

type TodoSeries struct {
	Series string
	Todos  []Todo
}

type Todo struct {
	ManifestTitle string `gorm:"type:varchar(255);"`
	ManifestMark  string `gorm:"type:varchar(500)"`
	ManifestTag   string `gorm:"type:varchar(1000);"`
	Series        string `gorm:"type:varchar(500);uniqueIndex"`
	TodoQuestions []TodoQuestion
}

type TodoQuestion struct {
	TodoId             uint
	Difficulty         string
	FrontendQuestionId string `gorm:"type:varchar(1000)"`
	Status             string `gorm:"type:varchar(500);commit:'enum Done|Doing|Null'"`
}

func (t *Todo) ConvFrom(tmpModel models.Todo) Todo {
	t.ManifestTag = tmpModel.ManifestTag
	t.ManifestMark = tmpModel.ManifestMark
	t.ManifestTitle = tmpModel.ManifestTitle
	t.Series = tmpModel.Series
	return *t
}

type AddTodo struct {
	Series       string
	ManifestList []string
}

func (a *AddTodo) Check() (bool, string) {

	if strings.TrimSpace(a.Series) == "" {
		return false, "未传入series"
	}

	if len(a.ManifestList) == 0 {
		return false, "add 参数未携带任何有效值"
	}

	return true, ""
}
