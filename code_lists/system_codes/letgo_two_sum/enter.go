package system_codes

import "letgo_repo/system_file/code_enter"

func init() {
	code_enter.Enter("system_codes", 1, twoSum, "[2,7,11,13],9")
}
