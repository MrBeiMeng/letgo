package letgo_two_sum_ii_input_array_is_sorted

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_twoSum(t *testing.T) {
	// 生成一个极端用例

	nums := make([]string, 0)

	target := 521

	leftBegin := -1000
	rightEnd := 1000

	for i := leftBegin; i <= rightEnd; i++ {
		if len(nums) == 0 {
			nums = append(nums, fmt.Sprintf("%d", i))
			continue
		}

		flag := false
		for _, value := range nums {
			num, _ := strconv.Atoi(value)
			if num+i == target {
				flag = true
				break
			}
		}
		if !flag {
			nums = append(nums, fmt.Sprintf("%d", i))
		}

		println(len(nums))

		if len(nums)%1000 == 0 {
			print("yes")
		}

		if len(nums) >= 29997 {

			fmt.Printf("数组长度:[%d]", len(nums))
			break
		}

	}

	fmt.Printf("[%v]", strings.Join(nums, ","))

}
