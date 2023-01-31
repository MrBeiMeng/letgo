package system_codes

/*罗马数字转整数 | https://leetcode.cn/problems/roman-to-integer*/

func romanToInt(s string) int {

	romanValueMap := getRomanValueMap()

	charByteArr := []byte(s)
	result := 0

	for index, charStr := range charByteArr {

		v := romanValueMap[charStr]

		if hasNext(index, charByteArr) && nextIsBigger(romanValueMap, charByteArr, index, v) {
			result -= v
			continue
		}

		result += v
	}

	return result
}

func getRomanValueMap() map[byte]int {
	romanValueMap := make(map[byte]int)
	romanValueMap['I'] = 1
	romanValueMap['V'] = 5
	romanValueMap['X'] = 10
	romanValueMap['L'] = 50
	romanValueMap['C'] = 100
	romanValueMap['D'] = 500
	romanValueMap['M'] = 1000
	return romanValueMap
}

func nextIsBigger(romanValueMap map[byte]int, charStrArr []byte, index int, v int) bool {
	return romanValueMap[charStrArr[index+1]] > v
}

func hasNext(index int, charArr []byte) bool {
	return index+1 < len(charArr)
}
