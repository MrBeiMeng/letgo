package letgo_maximum_repeating_substring

import (
	_ "letgo_repo/system_file/code_enter"
)

/*最大重复子字符串 | https://leetcode.cn/problems/maximum-repeating-substring*/
func maxRepeating(sequence string, word string) int {

	maxK := 0
	tmpK := 0
	for i := range sequence {
		for j := i; j < len(sequence); j++ {
			if j+len(word) > len(sequence) {
				tmpK = 0
				break
			}

			if sequence[j:j+len(word)] != word {
				tmpK = 0
				break
			}
			tmpK++
			j += len(word) - 1 // 因为j会加一
			maxK = max(maxK, tmpK)

			if j+1 == len(sequence) {
				return maxK
			}
		}
	}

	return maxK
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
