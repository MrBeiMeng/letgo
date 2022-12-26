package type_def

type CodeQueryWrapper struct {
	Level      string
	Star       string
	Status     string
	ShowHidden bool
	CodeNum    int
	CodeTitle  string
	Tags       []string
}

type Args struct {
	LinkedLists string
}

func (a *Args) IsEmpty() bool {
	if a.LinkedLists != "" {
		return false
	}

	return true
}
