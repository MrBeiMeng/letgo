package type_def

import (
	"letgo_repo/system_file/data_access/models"
	"strings"
)

type Manifest struct {
	QuestionsFrontIds []string
	Title             string
	Mark              string
	TagMap            map[string]struct{}
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
		m.TagMap[tag] = struct{}{}
	}
}

type Manifests []Manifest

func (m *Manifests) ConvFromModels(modelsData []models.Manifest) {
	for _, item := range modelsData {
		tmpManifest := Manifest{}
		tmpManifest.ConvFromModel(item)
		*m = append(*m, tmpManifest)
	}
}
