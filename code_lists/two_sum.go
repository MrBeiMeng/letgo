package code_lists

/*两数之和 | https://leetcode.cn/problems/two-sum*/

/*给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。
*/

func twoSum(nums []int, target int) []int {
	// 寻找的是下标数组
	indexMap := make(map[int]int)

	for i, num := range nums {
		if index, ok := indexMap[target-num]; ok {
			return []int{i, index}
		}

		indexMap[num] = i
	}

	return nil
}

type LetGoTwoSum struct { // todo 结构体名加上 LetGo前缀
}

func (p LetGoTwoSum) GetTest() string {
	return "[2,7,11,15],9"
}

func (p LetGoTwoSum) GetFunc() interface{} {
	return twoSum
}

func (p LetGoTwoSum) GetCodeNum() int {
	return 1
}

//func (p LetGoTwoSum) GetTags() []string {
//	return []string{enum.ARR, enum.HASH}
//}
//
//func (p LetGoTwoSum) RunDemo() {
//
//	p.twoSum(ArgsHandlerV1.getIntArr("[2,7,11,15]"), ArgsHandlerV1.getInt("9"))
//
//	arr := ArgsHandlerV1.getIntArr("[2,7,11,15]")
//	num := ArgsHandlerV1.getInt("9")
//	fmt.Printf("\tnums = %v\ttarget = %v\n", arr, num)
//
//	result := twoSum(arr, num)
//
//	fmt.Printf("\t%v", result)
//}
//

//
//func (p LetGoTwoSum) Run(args Args) {
//	// todo
//}
