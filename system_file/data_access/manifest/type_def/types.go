package type_def

type QueryWrapper struct {
	Title      string
	Tag        string
	TitleSlice []string
	TagSlice   []string
}

func (wrapper QueryWrapper) CaseNothing() bool {
	return wrapper.Title == "" && wrapper.Tag == "" && len(wrapper.TitleSlice) == 0 && len(wrapper.TagSlice) == 0
}
