package param_def

import (
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/utils/logger"
)

type ManifestCmdWrapper struct {
	Add    string
	Create bool
	Show   string
	Remove string
	Clear  bool
	Set    []string
}

func (m *ManifestCmdWrapper) CaseAdd() bool {
	return m.Add != ""
}

func (m *ManifestCmdWrapper) CaseShow() bool {
	return m.Show != ""
}

type TodoCmdWrapper struct {
	Series       string
	Default      bool
	CreateSeries string
	Add          []string
	Show         bool
	Remove       string
	Clear        bool
	Set          []string
}

func (m *TodoCmdWrapper) CaseAdd() bool {
	return len(m.Add) != 0
}

func (m *TodoCmdWrapper) CaseCreateSeries() bool {
	return m.CreateSeries != ""
}

func (m *TodoCmdWrapper) CaseChangeDefault() bool {
	return m.Default
}

func (m *TodoCmdWrapper) GetSeriesOrDefault() string {
	if m.Series == "" {
		defaultSeries, err := service.SGroupV1.ServiceTodo.GetDefaultSeriesName()
		if err != nil {
			logger.Logger.Break("必须携带--series 参数 或 设置默认系列")
		}
		logger.Logger.Info("使用默认系列 [%s]", defaultSeries)
		m.Series = defaultSeries
	}

	return m.Series
}

func (m *TodoCmdWrapper) CheckSeries() bool {

	hasSeries := m.Series != ""

	return hasSeries
}

type VersionWrapper struct {
	Detail bool
}
