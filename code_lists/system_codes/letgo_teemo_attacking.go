package system_codes

/*提莫攻击 | https://leetcode.cn/problems/teemo-attacking*/

func findPoisonedDuration(timeSeries []int, duration int) int {
	if duration == 0 {
		return 0
	}

	result := 0
	poisoningFlag := 0

	for _, t := range timeSeries {
		// 如果flag是0 或者到期 获取新的flag
		if poisoningFlag == 0 || t > poisoningFlag {
			poisoningFlag = getFlagUntil(t, duration)

			result += duration
			continue
		}

		result += getFlagUntil(t, duration) - poisoningFlag
		poisoningFlag = getFlagUntil(t, duration)
	}

	return result
}

func getFlagUntil(t int, duration int) int {
	return t + duration - 1
}
