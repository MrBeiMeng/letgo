package letgo_generate_parentheses

import "letgo_repo/system_file/code_enter"

func init() {
	code_enter.Enter("restart", 22, generateParenthesis)
}
