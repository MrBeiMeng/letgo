package utils

import (
	"fmt"
	"strings"
)

//func PrintLeetCodeProject(l type_def.LeetCodeProject) {
//	str := fmt.Sprintf("\t标题： %s\t标签： %s\t链接： %s\n\t等级： %s\t评价：%s\t状态：%s\n\t描述： %s\n\t示例： %s\n\t参数： %s",
//		l.GetCodeTitle(), strings.Join(l.GetTags(), "|"), l.GetUrl(), l.GetLevel(), l.GetStar(), l.GetStatus(), l.GetDescription(), l.GetExamples(), l.GetArgsDescription())
//
//	println(str)
//}
//
//func PrintLeetCodeProjectEasy(l type_def.LeetCodeProject) {
//	RowPrint("no.", "title", "level", "tags", "url")
//	RowPrint(l.GetCodeNum(), l.GetCodeTitle(), l.GetLevel(), strings.Join(l.GetTags(), "、"), l.GetUrl())
//	//str := fmt.Sprintf("\t|%d\t|%s\t|%s\t|[%s]\t|%s",
//	//	l.GetCodeNum(), l.GetCodeTitle(), l.GetLevel(), strings.Join(l.GetTags(), "、"), l.GetUrl())
//
//	//println(str)
//}

var GlobalRetraction = make([]string, 0)

func TablePrint(table [][]any) {
	if len(table) == 0 {
		return
	}

	// 统计每列最长的长度
	columnWList := make([]int, len(table[1]))
	for _, a := range table {
		if a == nil {
			continue
		}
		for j, item := range a {
			columnWidth := len(fmt.Sprintf("%v", item))
			if columnWidth >= columnWList[j] {
				columnWList[j] = columnWidth
			}
		}
	}

	// 补齐每列的长度之后再打印
	for _, row := range table {
		rowStrList := make([]any, 0)
		for j, item := range row {
			rowStrList = append(rowStrList, ConvLengthTo(item, columnWList[j]))
		}
		RowPrint(rowStrList...)
	}
}

func ConvLengthTo(a any, length int) string {
	str := fmt.Sprintf("%v", a)
	for len(str) < length {
		str += " "
	}

	return str
}

func RowPrint(a ...any) {
	for _, item := range a {
		fmt.Printf("\t|%v", item)
	}
	println()
}

// TPrint
//
//	@Description: 缩进打印
//	@param format
//	@param a
func TPrint(format string, a ...any) {
	GlobalRetraction = append(GlobalRetraction, "\t")

	if len(a) == 0 {
		fmt.Printf("%s%s\n", strings.Join(GlobalRetraction, ""), format)
		return
	}
	fmt.Printf("%s%v\n", strings.Join(GlobalRetraction, ""), fmt.Sprintf(format, a))
}

func TPrint2(format string, index int, a ...any) {
	GlobalRetraction = GlobalRetraction[:index]
	fmt.Printf("%s%v\n", GlobalRetraction, a)
}

// BPrint
//
//	@Description: 退格打印
//	@param a
func BPrint(a any) {
	GlobalRetraction = GlobalRetraction[:len(GlobalRetraction)-1]
	fmt.Printf("%s%v\n", GlobalRetraction, a)
}

func PrintArr(a ...any) {
	for _, item := range a {
		fmt.Printf("%v", item)
	}
	println()
}
