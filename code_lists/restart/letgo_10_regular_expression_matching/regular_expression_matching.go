package letgo_regular_expression_matching

import (
	_ "letgo_repo/system_file/code_enter"
)

/*正则表达式匹配 | https://leetcode.cn/problems/regular-expression-matching*/

/*
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。


示例 1：

输入：s = "aa", p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:

输入：s = "aa", p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：

输入：s = "ab", p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。


提示：

1 <= s.length <= 20
1 <= p.length <= 20
s 只包含从 a-z 的小写字母。
p 只包含从 a-z 的小写字母，以及字符 . 和 *。
保证每次出现字符 * 时，前面都匹配到有效的字符
*/

func isMatch(s string, p string) bool {
	// 使用正则表达式匹配 s 和 匹配模式 p

	matchDp := getDpPool(s, p)

	initDpByP(matchDp, p) // 初始化第一行dp，例如出现 a*b*c* 这种可有可无的情况

	// p=.*abc s=babababababbaabc
	// p=a*abc s=aaaabc

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if s[i] == p[j] || p[j] == '.' { // 相同
				matchDp[i+1][j+1] = matchDp[i][j]
				continue
			}

			if p[j] == '*' { // 如果*之前没有字符，这里会报错，但题中已有保证
				if p[j-1] == s[i] || p[j-1] == '.' { // 匹配模式前一位与当前s[i]匹配,也就是说s[i]可能重复1到多个。
					matchDp[i+1][j+1] = matchDp[i][j+1] || matchDp[i+1][j+1]
				}
				matchDp[i+1][j+1] = matchDp[i+1][j-1] || matchDp[i+1][j+1]

			}
		}
	}

	return matchDp[len(s)][len(p)]
}

func initDpByP(dp [][]bool, p string) {
	dp[0][0] = true

	for j := 0; j < len(p); j++ {
		if p[j] == '*' {
			dp[0][j+1] = dp[0][j-1]
		}
	}
}

// getDpPool
//
//	@Description: 生成记忆池
//	@param s
//	@param p
//	@return [][]bool
func getDpPool(s, p string) [][]bool {
	answer := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		answer[i] = make([]bool, len(p)+1)
	}
	return answer
}
