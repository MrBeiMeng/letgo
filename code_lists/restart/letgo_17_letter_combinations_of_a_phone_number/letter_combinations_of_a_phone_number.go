package letgo_letter_combinations_of_a_phone_number

import (
	_ "letgo_repo/system_file/code_enter"
)

/*电话号码的字母组合 | https://leetcode.cn/problems/letter-combinations-of-a-phone-number*/

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。


示例 1：

输入：digits = "23"
输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
示例 2：

输入：digits = ""
输出：[]
示例 3：

输入：digits = "2"
输出：["a","b","c"]


提示：

0 <= digits.length <= 4
digits[i] 是范围 ['2', '9'] 的一个数字。

*/

func letterCombinations(digits string) []string {
	alphabet2 := [][]byte{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}

	answer := make([]string, 0)

	backTrack([]byte(digits), 0, alphabet2, "", &answer)

	return answer
}

func backTrack(digits []byte, deep int, alphabet [][]byte, tmpStr string, answer *[]string) {
	if deep >= len(digits) {
		if tmpStr != "" {
			*answer = append(*answer, tmpStr)
		}
		return
	}

	b := digits[deep] - 48 - 2
	for _, al := range alphabet[b] {
		backTrack(digits, deep+1, alphabet, tmpStr+string(al), answer)
	}
}
