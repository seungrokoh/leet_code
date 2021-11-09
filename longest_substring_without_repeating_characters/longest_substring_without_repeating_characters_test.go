package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 0, lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	var start int
	var indexMap = make(map[byte]int)
	var max int
	length := len(s)
	for i := 0; i < length; i++ {
		index, ok := indexMap[s[i]]
		if ok {
			for j := start; j <= index; j++ {
				delete(indexMap, s[j])
			}
			start = index + 1
			indexMap[s[i]] = i
		} else {
			indexMap[s[i]] = i
			if max < len(indexMap) {
				max = len(indexMap)
			}
		}
	}
	return max
}
