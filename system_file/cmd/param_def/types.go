package param_def

type ManifestCmdWrapper struct {
	Add    string
	Create bool
	Show   bool
	Remove string
	Clear  bool
	Set    []string
}

func (m *ManifestCmdWrapper) CaseAdd() bool {
	return m.Add != ""
}

func (m *ManifestCmdWrapper) CaseShow() bool {
	return m.Show
}

type TodoCmdWrapper struct {
	Series       string
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
