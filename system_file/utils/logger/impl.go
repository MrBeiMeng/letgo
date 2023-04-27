package logger

import (
	"fmt"
	"letgo_repo/system_file/utils"
	"os"
)

type LoggerImpl struct {
}

func (l LoggerImpl) Success(msg string, asg ...interface{}) {
	fullMsg := getMsg(msg, asg...)
	println(utils.GetColorGreen(fmt.Sprintf("info: %s", fullMsg)))
}

func (l LoggerImpl) Info(msg string, asg ...interface{}) {
	fullMsg := getMsg(msg, asg...)
	println(utils.GetColorWhite(fmt.Sprintf("info: %s", fullMsg)))
}

func (l LoggerImpl) Warn(msg string, asg ...interface{}) {
	fullMsg := getMsg(msg, asg...)
	println(utils.GetColorPurple(fmt.Sprintf("warn: %s", fullMsg)))
}

func (l LoggerImpl) Break(msg string, asg ...interface{}) {
	fullMsg := getMsg(msg, asg...)
	println(utils.GetColorRed(fmt.Sprintf("error: %s", fullMsg)))
	os.Exit(0)
}

func getMsg(msg string, asg ...interface{}) string {
	return fmt.Sprintf(msg, asg...)
}
