package common

import (
	"letgo_repo/system_file/service/type_def"
	"letgo_repo/system_file/utils"
	"strings"
)

type ServiceImpl struct {
}

func (s *ServiceImpl) GetVersionHistories() []string {
	answer := make([]string, 0)

	var versionBody type_def.VersionBody

	err := versionBody.InitByJsonFile("version.json")
	if err != nil {
		panic(err)
	}

	// 读取version.json 文件.
	detailFormatTemplate := "${project_name} v${version_no} -- ${type}[${date}] | ${log}"

	tmpStr := strings.Builder{}

	for i, perVersion := range versionBody.Versions {
		argsMap := make(map[string]string)
		argsMap["project_name"] = versionBody.ProjectName
		argsMap["version_no"] = perVersion.VersionNo
		argsMap["type"] = perVersion.Type
		argsMap["date"] = perVersion.Date
		argsMap["log"] = perVersion.Log

		if i == 0 {
			tmpStr.WriteString("🔎 ")
		} else {
			tmpStr.WriteString("   ")
		}

		tmpStr.WriteString(utils.ReplaceAll(detailFormatTemplate, argsMap))

		if i == 0 {
			tmpStr.WriteString(" ✍️")
		}

		answer = append(answer, tmpStr.String())
		tmpStr.Reset()
	}

	return answer
}
