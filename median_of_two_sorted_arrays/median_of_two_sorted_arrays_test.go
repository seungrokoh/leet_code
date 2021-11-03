package median_of_two_sorted_arrays

import (
	"testing"
)

func TestSortedArrays(t *testing.T) {
	var nums1 = []int{1, 2}
	var nums2 = []int{3, 4}
	findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	var sortedArray []int
	var offset1 int
	var offset2 int
	for offset1 < len1 && offset2 < len2 {
		if nums1[offset1] < nums2[offset2] {
			sortedArray = append(sortedArray, nums1[offset1])
			offset1++
		} else if nums1[offset1] > nums2[offset2] {
			sortedArray = append(sortedArray, nums2[offset2])
			offset2++
		} else {
			sortedArray = append(sortedArray, nums1[offset1], nums2[offset2])
			offset1++
			offset2++
		}
	}
	for i := offset1; i < len1; i++ {
		sortedArray = append(sortedArray, nums1[i])
	}
	for i := offset2; i < len2; i++ {
		sortedArray = append(sortedArray, nums2[i])
	}
	middle := (len1 + len2) / 2
	if len(sortedArray)%2 == 0 {
		return float64(sortedArray[middle]+sortedArray[middle-1]) / 2
	}
	return float64(sortedArray[middle])
}
