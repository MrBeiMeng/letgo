package utils

import (
	"fmt"
	"os"
	"strings"
)

func ConvLineToCamel(name string) (camel string) {
	strs := strings.Split(name, "-")

	for i, str := range strs {
		bytes := []byte(str)
		bytes[0] = str[0] - 32
		strs[i] = string(bytes)
	}
	return strings.Join(strs, "")
}

func InitFile(slug, url, title string, codeNum int) {
	file, err := os.Create(fmt.Sprintf("code_lists/%s.go", strings.ReplaceAll(slug, "-", "_")))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	template := `package code_lists
/*${title} | ${url}*/

type ${structName} struct {
}

func (p ${structName}) GetTags() []string {
	return []string{} // todo 标签
}

func (p ${structName}) RunDemo() {
	// todo
}

func (p ${structName}) GetCodeNum() int {
	return ${codeNum}
}

func (p ${structName}) Run(args Args) {
	// todo
}
`
	template = strings.ReplaceAll(template, "${structName}", ConvLineToCamel(slug))
	template = strings.ReplaceAll(template, "${codeNum}", fmt.Sprintf("%d", codeNum))
	template = strings.ReplaceAll(template, "${url}", fmt.Sprintf("%s", url))
	template = strings.ReplaceAll(template, "${title}", fmt.Sprintf("%s", title))

	file.WriteString(template)
}
