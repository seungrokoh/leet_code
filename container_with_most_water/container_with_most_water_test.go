package container_with_most_water

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	assert.Equal(t, 49, maxArea(height))
}

func maxArea(height []int) int {
	length := len(height)
	left := 0
	right := length - 1

	var max int
	for left < right {
		width := right - left
		r := min(height[left], height[right]) * width
		if max < r {
			max = r
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
