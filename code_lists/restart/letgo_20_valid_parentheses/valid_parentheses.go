package letgo_valid_parentheses

import (
	_ "letgo_repo/system_file/code_enter"
)

/*有效的括号 | https://leetcode.cn/problems/valid-parentheses*/

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"
输出：true
示例   2：

输入：s = "()[]{}"
输出：true
示例   3：

输入：s = "(]"
输出：false

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/

type Stack struct {
	arr []byte
}

func (s *Stack) Push(val byte) {
	s.arr = append(s.arr, val)
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Pop() (answer byte, ok bool) {
	if len(s.arr) > 0 {
		answer = s.arr[len(s.arr)-1]
		ok = true
		s.arr = s.arr[:len(s.arr)-1]
	}

	return answer, ok
}

func isValid(s string) bool {
	stack := Stack{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '}':
			if val, ok := stack.Pop(); !ok || val != '{' {
				return false
			}
		case ']':
			if val, ok := stack.Pop(); !ok || val != '[' {
				return false
			}
		case ')':
			if val, ok := stack.Pop(); !ok || val != '(' {
				return false
			}
		default:
			stack.Push(s[i])
		}
	}

	return stack.IsEmpty()
}
