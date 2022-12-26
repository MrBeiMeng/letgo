package utils

import "fmt"

func GetColorYellow(inject string) string {
	return fmt.Sprintf("\033[1;33m%s\u001B[0m\t", inject)
}

func GetColorWhite(inject string) string {
	return fmt.Sprintf("\033[1;37m%s\u001B[0m\t", inject)
}

func GetColorGreen(inject string) string {
	return fmt.Sprintf("\033[1;32m%s\u001B[0m\t", inject)
}

func GetColorBlue(inject string) string {
	return fmt.Sprintf("\033[1;34m%s\u001B[0m\t", inject)
}

func GetColorRed(inject string) string {
	return fmt.Sprintf("\033[1;35m%s\u001B[0m\t", inject)
}

func GetEasy() string {
	return GetColorGreen("简单")
}

func GetMid() string {
	return GetColorBlue("中等")
}

func GetHard() string {
	return GetColorRed("困难")
}

// GetStar 1 到 5
func GetStar(point int) string {
	str := ""
	for i := 0; i < point; i++ {
		str += "⭐️"
	}
	return GetColorYellow(str)
}

func GetDoing() string {
	return ""
}

func GetDone() string {
	return GetColorWhite("DONE")
}
