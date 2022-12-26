package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetInput(message string, level int) (string, error) {
	println(message)
	print("bl>")
	for i := 0; i < level; i++ {
		print(">")
	}
	inputStr := ""
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}
	inputStr = string(line)

	if strings.EqualFold(inputStr, "quit") || strings.EqualFold(inputStr, "stop") {
		return "", errors.New("用户结束进程")
	}

	return inputStr, nil
}
