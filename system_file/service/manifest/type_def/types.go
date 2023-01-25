package type_def

import (
	"fmt"
	"letgo_repo/system_file/data_access/models"
	"letgo_repo/system_file/utils"
	"strings"
)

type Manifest struct {
	QuestionsFrontIds []string
	Title             string
	Mark              string
	TagMap            map[string]struct{}
}

func (m *Manifest) GetTags() string {
	builder := strings.Builder{}
	for key, _ := range m.TagMap {
		if builder.Len() == 0 {
			builder.WriteString(key)
			continue
		}

		builder.WriteString(fmt.Sprintf(",%s", key))
	}

	return builder.String()
}

func (m *Manifest) AppendTag(tags ...string) {
	for _, tag := range tags {
		if m.TagMap == nil {
			m.TagMap = make(map[string]struct{})
		}
		m.TagMap[tag] = struct{}{}
	}

	return
}

func (m *Manifest) ConvToModel() (tmpObj models.Manifest) {
	tmpObj.Title = m.Title

	arr := make([]string, 0)
	for key, _ := range m.TagMap {
		arr = append(arr, key)
	}

	tmpObj.Tags = strings.Join(arr, ",")
	tmpObj.QuestionsFrontIds = strings.Join(m.QuestionsFrontIds, ",")
	tmpObj.Mark = m.Mark

	return
}

func (m *Manifest) ConvFromModel(tmpObj models.Manifest) {
	m.Title = tmpObj.Title
	m.Mark = tmpObj.Mark
	m.QuestionsFrontIds = strings.Split(tmpObj.QuestionsFrontIds, ",")
	for _, tag := range strings.Split(tmpObj.Tags, ",") {
		if m.TagMap == nil {
			m.TagMap = make(map[string]struct{})
		}

		m.TagMap[tag] = struct{}{}
	}
}

type ManifestWithDetail struct {
	Manifest
	Questions []models.Question
}

func (m *ManifestWithDetail) GetFrontIdsWithColor() string {
	builder := strings.Builder{}

	for _, question := range m.Questions {

		codeNum := question.FrontendQuestionId

		switch question.Difficulty {
		case "EASY":
			codeNum += utils.GetColorCyan("·")
		case "MEDIUM":
			codeNum += utils.GetColorYellow("·")
		case "HARD":
			codeNum += utils.GetColorPurple("·")
		}

		if builder.Len() != 0 {
			codeNum = "," + codeNum
		}

		builder.WriteString(codeNum)
	}

	return builder.String()
}

type Manifests []ManifestWithDetail

func (m *Manifests) ConvFromModels(modelsData []models.Manifest) {
	for _, item := range modelsData {
		tmpManifest := ManifestWithDetail{}
		tmpManifest.ConvFromModel(item)
		*m = append(*m, tmpManifest)
	}
}
