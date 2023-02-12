package letgo_teemo_attacking

import (
	_ "letgo_repo/system_file/code_enter"
)

/*提莫攻击 | https://leetcode.cn/problems/teemo-attacking*/

func findPoisonedDuration(timeSeries []int, duration int) int {
	until := -1
	ans := 0
	for _, ts := range timeSeries {
		if ts <= until { // [... now <= until ... newUntil .]
			ans += getUntil(ts, duration) - until
			until = getUntil(ts, duration)
			continue
		}

		until = getUntil(ts, duration)
		ans += duration
	}

	return ans
}

func getUntil(ts int, duration int) int { // 获取新的 毒性消除边界
	return ts + duration - 1
}
