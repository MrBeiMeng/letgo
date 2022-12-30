package main

import (
	"fmt"
	"letgo_repo/data_access"
	"letgo_repo/utils"
)

func main() {
	err := data_access.GenerateTable()
	if err != nil {
		println(utils.GetColorRed(fmt.Sprintf("创建失败%s", err.Error())))
		return
	}

	fmt.Printf(utils.GetColorGreen("创建成功"))
}
