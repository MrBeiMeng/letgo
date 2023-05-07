package letgo_longest_palindromic_substring

import (
	_ "letgo_repo/system_file/code_enter"
)

/*最长回文子串 | https://leetcode.cn/problems/longest-palindromic-substring*/

/*
给你一个字符串 s，找到 s 中最长的回文子串。

如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"

提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成

*/

func longestPalindrome(s string) string {
	// 动态规划

	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}

	dpAnswer, index1, index2 := 0, 0, 0

	// 搞一个dp数组[i][j]，i 表示字符串起点，j 表示字符串终点 dp[i][j]表示 从i到j是否为回文字符串
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j <= len(s)-1; j++ {
			if i == j { // 对角线处默认为回文字符串，因为一个字母也是回文字符串
				dp[i][j] = 1
				continue
			}

			if s[i] != s[j] { // 题中规定只会出现英文和数组，ascii码足以应对了
				dp[i][j] = 0
				continue
			}

			if isPalindrome(dp, i, j) { // 两个点相同且位置不同的情况下

				dp[i][j] = dp[i+1][j-1] + 2

				if dp[i][j] >= dpAnswer {
					dpAnswer = dp[i][j]
					index1 = i
					index2 = j
				}
			}
		}
	}

	return s[index1 : index2+1]
}

// isPalindrome
//
//	@Description: 是否是回文字符串
//	@param dp
//	@param i
//	@param j
//	@return bool
func isPalindrome(dp [][]int, i int, j int) bool {
	return dp[i+1][j-1] > 0 || j-i == 1
}
