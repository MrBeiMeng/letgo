package letgo_string_matching_in_an_array

import (
	_ "letgo_repo/system_file/code_enter"
)

/*数组中的字符串匹配 | https://leetcode.cn/problems/string-matching-in-an-array*/

func stringMatching(words []string) []string {
	answer := make([]string, 0)
	for i, pattern := range words {
		for j, str := range words {
			if i == j {
				continue
			}

			if kmpSearch(str, pattern) > -1 {
				answer = append(answer, pattern)
				break
			}
		}

	}

	return answer
}

func kmpSearch(str, pattern string) int {
	next := kmpNext(pattern)

	nIndex := 0
	for i := range str {
		if nIndex > 0 && str[i] != pattern[nIndex] {
			nIndex = next[nIndex-1]
		}

		if str[i] == pattern[nIndex] {
			nIndex++
		}

		if nIndex == len(next) {
			return i - len(next) + 1
		}
	}

	return -1
}

func kmpNext(pattern string) []int {
	next := make([]int, len(pattern))

	prefixIndex := 0
	for i := 1; i < len(pattern); i++ {
		if prefixIndex > 0 && pattern[i] != pattern[prefixIndex] {
			prefixIndex = next[prefixIndex-1]
		}

		if pattern[i] == pattern[prefixIndex] {
			prefixIndex++
			next[i] = prefixIndex
		} else {
			next[i] = 0
		}
	}

	return next
}
