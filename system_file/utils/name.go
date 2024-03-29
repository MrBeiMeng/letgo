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

func InitFile(series, slug, url, titleCn string, codeNum int, code string) {
	fileName := strings.ReplaceAll(slug, "-", "_")

	dirPath := fmt.Sprintf("code_lists/%s/letgo_%d_%s", series, codeNum, fileName)
	os.MkdirAll(dirPath, 0777)

	file, err := os.Create(fmt.Sprintf("%s/%s.go", dirPath, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	template := `package ${package}

import (
	_ "letgo_repo/system_file/code_enter"
)

/*${title} | ${url}*/

${code}
`
	template = strings.ReplaceAll(template, "${package}", fmt.Sprintf("letgo_%s", fileName))
	template = strings.ReplaceAll(template, "${title}", titleCn)
	template = strings.ReplaceAll(template, "${code}", fmt.Sprintf("%s", code))
	template = strings.ReplaceAll(template, "${url}", fmt.Sprintf("%s", url))

	template = strings.ReplaceAll(template, "{\n\n}", "{\n\t//TODO implement me\n\tpanic(\"implement me\")\n}")

	file.WriteString(template)

	file2, err := os.Create(fmt.Sprintf("%s/enter.go", dirPath))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	template2 := `package ${package}

import "letgo_repo/system_file/code_enter"

func init() {
	code_enter.Enter("${series}", ${codeNum}, ${methodName})
}
`
	methodName := strings.Split(strings.Split(code, "(")[0], "func ")[1]

	template2 = strings.ReplaceAll(template2, "${package}", fmt.Sprintf("letgo_%s", fileName))
	template2 = strings.ReplaceAll(template2, "${series}", series)
	template2 = strings.ReplaceAll(template2, "${codeNum}", fmt.Sprintf("%d", codeNum))
	template2 = strings.ReplaceAll(template2, "${methodName}", methodName)

	file2.WriteString(template2)

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
	oFile, err := os.OpenFile(fmt.Sprintf("code_lists/%s/enter.go", series), syscall.O_CREAT, 666)
	if err != nil {
		panic(err)
	}

	all, err := ioutil.ReadAll(oFile)
	if err != nil {
		panic(err)
	}
	oFile.Close()

	newLine := fmt.Sprintf("\"\n\t_ \"letgo_repo/%s\"\n)", dirPath)

	allStr := ""
	if strings.Contains(string(all), "import") {
		//allStr := strings.ReplaceAll(string(all), "// import at here", strings.ReplaceAll(newLine, "${structName}", ConvLineToCamel(slug)))
		allStr = strings.ReplaceAll(string(all), "\"\n)", newLine)
	} else if strings.Contains(string(all), "package") {
		allStr = strings.ReplaceAll(string(all), fmt.Sprintf("package %s\n", series), fmt.Sprintf(`package %s
import (
	_ "letgo_repo/%s"
)`, series, dirPath))
	} else {
		allStr = fmt.Sprintf(`package %s
import (
	_ "letgo_repo/%s"
)

var series string = "%s"

func init() {
	// enter new code here
}`, series, dirPath, series)
	}

	bFile, err := os.OpenFile(fmt.Sprintf("code_lists/%s/enter.go", series), syscall.O_RDWR, 777)
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
	//
	//// 设定一个首层切割标志,当面一旦超过首层便不再切割了。
	//fistSlice := true
	//builder := strings.Builder{}
	//
	//for _, char := range s {
	//	if !fistSlice {
	//		builder.WriteRune(char)
	//		continue
	//	}
	//
	//
	//
	//
	//}

	str := ""
	splitFlag := 0
	for _, char := range s {
		if char == '[' {
			splitFlag += 1
		}

		if splitFlag != 0 {
			str = fmt.Sprintf("%s%c", str, char)
			if char == ']' {
				splitFlag -= 1
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
