package system_codes

import (
	"strconv"
)

/*Fizz Buzz | https://leetcode.cn/problems/fizz-buzz*/

// fizzBuzz
//
//	@Description: answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
//
// answer[i] == "Fizz" 如果 i 是 3 的倍数。
// answer[i] == "Buzz" 如果 i 是 5 的倍数。
// answer[i] == i （以字符串形式）如果上述条件全不满足。
//
//	@param n
//	@return []string
func fizzBuzz(n int) []string {

	resultList := make([]string, n)

	for i := range resultList {
		num := i + 1
		if num%3 != 0 && num%5 != 0 {
			resultList[i] = strconv.Itoa(num)
			continue
		}

		if num%3 == 0 {
			resultList[i] += "Fizz"
		}

		if num%5 == 0 {
			resultList[i] += "Buzz"
		}
	}

	return resultList
}
