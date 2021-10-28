package longest_common_prefix

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	var minIndex int
	var minLen = math.MaxInt8
	for i, s := range strs {
		if minLen > len(s) {
			minIndex = i
			minLen = len(s)
		}
	}

	target := strs[minIndex]
	var result string
	for i := range target {
		for _, s := range strs {
			if s[i] != target[i] {
				return result
			}
		}
		result = fmt.Sprintf("%s%c", result, target[i])
	}
	return result
}
