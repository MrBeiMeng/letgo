package letgo_generate_parentheses

import (
	_ "letgo_repo/system_file/code_enter"
)

/*括号生成 | https://leetcode.cn/problems/generate-parentheses*/

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]
示例 2：
输入：n = 1
输出：["()"]

提示：

1 <= n <= 8
*/

func generateParenthesis(n int) []string {
	answer := make([]string, 0)
	backTrack("", 0, 0, n, &answer)
	return answer
}

func backTrack(str string, left, right, n int, answer *[]string) {
	// 如果左侧括号和右侧括号的数量一致并且都等于n，则完成此次递归,保存数据
	if left == right && left == n {
		*answer = append(*answer, str)
		return
	}

	// 如果左侧括号数量不够n，走左侧递归
	if left < n {
		backTrack(str+"(", left+1, right, n, answer)
	}

	if right < left {
		// 走右侧括号递归
		backTrack(str+")", left, right+1, n, answer)
	}
}
