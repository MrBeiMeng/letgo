package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"syscall"
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

	/*// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // 只读模式
	O_WRONLY int = syscall.O_WRONLY //只写模式
	O_RDWR   int = syscall.O_RDWR   // 读写混合模式
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // 写模式的时候将数据附加到文件末尾
	O_CREATE int = syscall.O_CREAT  // 文件如果不存在就新建
	O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE模式一起使用, 文件必须不存在
	O_SYNC   int = syscall.O_SYNC   //打开文件用于同步 I/O.
	O_TRUNC  int = syscall.O_TRUNC  // 打开文件时清空文件*/
	oFile, err := os.OpenFile("code_lists/enter.go", syscall.O_RDONLY, 666)
	if err != nil {
		panic(err)
	}

	all, err := ioutil.ReadAll(oFile)
	if err != nil {
		panic(err)
	}
	oFile.Close()

	newLine := `CodeChallengeList = append(CodeChallengeList, enterCodeChallenge(${structName}{}))
	// enter new code here`
	allStr := strings.ReplaceAll(string(all), "// enter new code here", strings.ReplaceAll(newLine, "${structName}", ConvLineToCamel(slug)))

	bFile, err := os.OpenFile("code_lists/enter.go", syscall.O_RDWR, 777)
	if err != nil {
		panic(err)
	}
	bFile.WriteString(allStr)

	bFile.Close()
}

// RoughSplit
//
//	@Description: 粗切割字符串 例:"[1,2],3" - ["1,2","3"]
//	@param s
//	@return []string
func RoughSplit(s string) (result []string) {
	// [1,2,3],4,5,[6,7],8,[9],[10]

	str := ""
	splitFlag := true
	for _, char := range s {
		if char == '[' {
			splitFlag = false
		}

		if !splitFlag {
			str = fmt.Sprintf("%s%c", str, char)
			if char == ']' {
				splitFlag = true
			}
			continue
		}

		if char == ',' {
			result = append(result, str)
			str = ""
			continue
		}

		str = fmt.Sprintf("%s%c", str, char)
	}

	if str != "" {
		result = append(result, str)
	}

	return result
}
