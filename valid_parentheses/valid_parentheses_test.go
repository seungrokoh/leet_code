package valid_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	var s = "()[]{}"
	assert.True(t, isValid(s))
	s = "([)]"
	assert.False(t, isValid(s))
	s = "("
	assert.False(t, isValid(s))
}

func isValid(s string) bool {
	var stack []byte

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, byte(c))
		default:
			last := len(stack) - 1
			if last < 0 {
				return false
			}
			switch c {
			case ')':
				if stack[last] != '(' {
					return false
				}
				stack = stack[:last]
			case '}':
				if stack[last] != '{' {
					return false
				}
				stack = stack[:last]
			case ']':
				if stack[last] != '[' {
					return false
				}
				stack = stack[:last]
			}
		}
	}
	return len(stack) == 0
}
