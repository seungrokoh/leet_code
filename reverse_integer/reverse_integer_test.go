package reverse_integer

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	var num = -123
	log.Println(reverse(num))
	assert.Equal(t, -321, reverse(num))
	num = 123
	assert.Equal(t, 321, reverse(num))
	num = 123000
	assert.Equal(t, 321, reverse(num))
}

func reverse(x int) int {
	oper := 1
	num := x
	for num != 0 {
		oper *= 10
		num /= 10
	}
	oper /= 10

	num = x
	var result int
	for num != 0 {
		result += (num % 10) * oper
		oper /= 10
		num /= 10
	}

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
