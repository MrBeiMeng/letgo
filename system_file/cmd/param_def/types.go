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
