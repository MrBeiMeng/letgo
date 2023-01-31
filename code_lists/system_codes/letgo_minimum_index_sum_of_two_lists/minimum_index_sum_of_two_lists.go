package letgo_minimum_index_sum_of_two_lists

import (
	_ "letgo_repo/system_file/code_enter"
)

/*两个列表的最小索引总和 | https://leetcode.cn/problems/minimum-index-sum-of-two-lists*/

func findRestaurant(list1 []string, list2 []string) []string {

	strMap := make(map[string]int)

	minimumSum := -1

	for i, str := range list1 {
		strMap[str] = i
	}

	answer := make([]string, 0)

	for i, str := range list2 {
		if i2, ok := strMap[str]; ok {
			if minimumSum != -1 && i2+i > minimumSum {
				continue
			}

			if i+i2 < minimumSum || minimumSum == -1 {
				minimumSum = i + i2
				answer = nil
			}

			answer = append(answer, str)
		}
	}

	return answer
}
