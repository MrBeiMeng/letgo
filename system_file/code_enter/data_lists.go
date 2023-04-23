package code_enter

import (
	"letgo_repo/system_file/data_access/todo"
	"letgo_repo/system_file/utils/logger"
)

type QuestionSolution struct {
	CodeNum int         // 题号
	RunFunc interface{} // 解法
	Tests   []string    //测试案例
}

type QuestionSolutions []QuestionSolution

var QuestionSolutionsV1 QuestionSolutions

var series string

func init() {

	series1, err := todo.DATodoV1.SelectDefaultSeriesName()
	if err != nil {
		//panic(err)
		logger.Logger.Warn(err.Error())
	} else {
		logger.Logger.Info("注入默认系列[%s]的解题代码", series1)
	}

	series = series1
}

func Enter(tmpSeries string, codeNum int, funci interface{}, tests ...string) {

	if series == tmpSeries {
		QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(codeNum, funci, tests...))
	}

	return
}

func GetProblemSolution(codeNum int, runFunc interface{}, tests ...string) (obj QuestionSolution) {
	obj.CodeNum = codeNum
	obj.RunFunc = runFunc
	obj.Tests = tests
	return obj
}
