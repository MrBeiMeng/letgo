package code_lists

type QuestionSolution struct {
	CodeNum int         // 题号
	RunFunc interface{} // 解法
	Tests   []string    //测试案例
}

type QuestionSolutions []QuestionSolution

var QuestionSolutionsV1 QuestionSolutions

func init() {
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(1, twoSum, "[2,7,11,13],9"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(234, isPalindrome, "[1,2,3,4,3,2,1]"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(485, findMaxConsecutiveOnes, "[1,1,0,1,1,1]", "[1,0,1,1,0,1]"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(495, findPoisonedDuration, "[1,4],2", "[1,2],2"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(13, romanToInt, "III", "IV", "IX", "LVIII", "MCMXCIV"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(383, canConstruct, ""))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(412, fizzBuzz, "3", "5", "15"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(876, middleNode, "[1,2,3,4,5]", "[1,2,3,4,5,6]"))
	QuestionSolutionsV1 = append(QuestionSolutionsV1, GetProblemSolution(1342, numberOfSteps, "14", "8", "123"))
	// enter new code here
}

func GetProblemSolution(codeNum int, runFunc interface{}, tests ...string) (obj QuestionSolution) {
	obj.CodeNum = codeNum
	obj.RunFunc = runFunc
	obj.Tests = tests
	return obj
}
