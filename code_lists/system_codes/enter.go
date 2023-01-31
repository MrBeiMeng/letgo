package system_codes

import (
	_ "letgo_repo/code_lists/system_codes/letgo_contains_duplicate_ii"
	_ "letgo_repo/code_lists/system_codes/letgo_two_sum"
	_ "letgo_repo/code_lists/system_codes/letgo_two_sum_ii_input_array_is_sorted"
	// import at here
	"letgo_repo/system_file/code_enter"
)

var series string = "system_codes"

func init() {
	code_enter.Enter(series, 234, isPalindrome, "[1,2,3,4,3,2,1]")
	code_enter.Enter(series, 485, findMaxConsecutiveOnes, "[1,1,0,1,1,1]", "[1,0,1,1,0,1]")
	code_enter.Enter(series, 495, findPoisonedDuration, "[1,4],2", "[1,2],2")
	code_enter.Enter(series, 13, romanToInt, "III", "IV", "IX", "LVIII", "MCMXCIV")
	code_enter.Enter(series, 383, canConstruct, "")
	code_enter.Enter(series, 412, fizzBuzz, "3", "5", "15")
	code_enter.Enter(series, 876, middleNode, "[1,2,3,4,5]", "[1,2,3,4,5,6]")
	code_enter.Enter(series, 1342, numberOfSteps, "14", "8", "123")
	code_enter.Enter(series, 1672, maximumWealth, "[[1,2,3],[3,2,1]]", "[[1,5],[7,3],[3,5]]", "[[2,8,7],[7,1,3],[1,9,5]]")
	code_enter.Enter(series, 414, thirdMax, "[3,2,1]", "[1,2]", "[2,2,3,1]")
	code_enter.Enter(series, 628, maximumProduct, "[1,2,3]", "[1,2,3,4]", "[-1,-2,-3]", "[-100,-98,-1,2,3,4]")
	code_enter.Enter(series, 645, findErrorNums, "[1,2,3,4,5,2]")
	code_enter.Enter(series, 697, findShortestSubArray)
	code_enter.Enter(series, 448, findDisappearedNumbers)
	code_enter.Enter(series, 442, findDuplicates)
	code_enter.Enter(series, 274, hIndex)
	code_enter.Enter(series, 41, firstMissingPositive)
	code_enter.Enter(series, 283, moveZeroes)
	code_enter.Enter(series, 453, minMoves)
	code_enter.Enter(series, 665, checkPossibility)
	code_enter.Enter(series, 118, generate)
	code_enter.Enter(series, 119, getRow)
	code_enter.Enter(series, 661, imageSmoother)
	code_enter.Enter(series, 598, maxCount)
	// enter new code here
}
