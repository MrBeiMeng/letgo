package utils

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strings"
)

var GlobalRetraction = make([]string, 0)

func TablePrint(strTable [][]string) {

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
	table.CloseBorder()
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
