package letgo_check_if_a_word_occurs_as_a_prefix_of_any_word_in_a_sentence

import (
	_ "letgo_repo/system_file/code_enter"
	"strings"
)

/*检查单词是否为句中其他单词的前缀 | https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence*/

func isPrefixOfWord(sentence string, searchWord string) int {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		if len(word) < len(searchWord) {
			continue
		}
		isPrefix := true

		for j := range searchWord {
			if word[j] != searchWord[j] {
				isPrefix = false
				break
			}
		}
		if isPrefix {
			return i + 1
		}
	}

	return -1
}
