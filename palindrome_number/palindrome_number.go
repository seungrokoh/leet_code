package palindrome_number

import "strconv"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 || x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	length := len(s)
	for i := 0; i <= length/2; i++ {
		if s[i] != s[length-1-i] {
			return false
		}
	}
	return true
}
