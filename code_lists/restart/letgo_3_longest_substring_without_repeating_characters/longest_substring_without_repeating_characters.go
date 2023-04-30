package letgo_longest_substring_without_repeating_characters

import (
	_ "letgo_repo/system_file/code_enter"
)

/*无重复字符的最长子串 | https://leetcode.cn/problems/longest-substring-without-repeating-characters*/

/*
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

tips:
	0 <= s.length <= 5 * 104
	s 由英文字母、数字、符号和空格组成
*/

func lengthOfLongestSubstring(s string) int {

	indexMap := make(map[byte]int)

	preIndex, answer := -1, 0

	for i := 0; i < len(s); i++ {

		if index, ok := indexMap[s[i]]; ok && index > preIndex {
			preIndex = index
		}

		indexMap[s[i]] = i
		answer = max(answer, i-preIndex)
	}

	return answer
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
