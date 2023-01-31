package system_codes

/*赎金信 | https://leetcode.cn/problems/ransom-note*/

// canConstruct
//
//	@Description: 判断ransomNote能否由 magazine 组成
//	@param ransomNote
//	@param magazine
//	@return bool
func canConstruct(ransomNote string, magazine string) bool {

	alphaList := make([]int, 26)

	for _, alpha := range []byte(magazine) {
		alphaList[alpha-97] += 1
	}

	for _, alpha := range []byte(ransomNote) {
		if alphaList[alpha-97] == 0 {
			return false
		}

		alphaList[alpha-97] -= 1
	}

	return true
}
