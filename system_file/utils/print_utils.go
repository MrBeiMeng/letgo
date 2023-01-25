package utils

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strings"
)

var GlobalRetraction = make([]string, 0)

func Max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func ConvStrLengthTo(str string, length int) string {
	num := length - GetStrLength(str)
	for i := 0; i < num; i++ {
		str += GetColorBlue("")
	}
	return str
}

func GetStrLength(str string) int {
	num := 0
	str = strings.ReplaceAll(str, "\u001B[0m", " ")

	num += strings.Count(str, "\033[1;00m")
	num += strings.Count(str, "\033[1;31m")
	num += strings.Count(str, "\033[1;32m")
	num += strings.Count(str, "\033[1;33m")
	num += strings.Count(str, "\033[1;34m")
	num += strings.Count(str, "\033[1;35m")
	num += strings.Count(str, "\033[1;36m")
	num += strings.Count(str, "\033[1;37m")

	return num
}

func TablePrintColorHandler(strTable [][]string, colorColumn []int) {
	lengthArr := make([]int, len(strTable[0]))

	for _, arr := range strTable {
		for _, index := range colorColumn {
			lengthArr[index] = Max(GetStrLength(arr[index]), lengthArr[index])
		}
	}

	for i := 0; i < len(strTable); i++ {
		for _, index := range colorColumn {
			strTable[i][index] = ConvStrLengthTo(strTable[i][index], lengthArr[index])
		}
	}

	TablePrint(strTable, false)
}

func TablePrint(strTable [][]string, borderVisible bool) {

	table, err := gotable.Create(strTable[0]...)
	if err != nil {
		println(err.Error())
		return
	}

	for i := 1; i < len(strTable); i++ {
		err = table.AddRow(strTable[i])
		if err != nil {
			println(err.Error())
		}
	}
	if !borderVisible {
		table.CloseBorder()
	}
	println(table.String())
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
