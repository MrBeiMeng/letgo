package code_lists

/*将数字变成 0 的操作次数 | https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-to-zero*/

// numberOfSteps
//
//	@Description: 给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。
//	@param num
//	@return int
func numberOfSteps(num int) (count int) {

	for num != 0 {
		if num-1 == 0 {
			return count + 1
		}

		count = count + 1 + num%2
		num = num >> 1
	}

	return
}
