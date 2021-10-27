package two_sum

func twoSum(nums []int, target int) []int {
	data := make(map[int]int, len(nums))
	for index, num := range nums {
		data[num] = index
	}
	var result []int
	for index, num := range nums {
		if i, ok := data[target-num]; (index != i) && ok {
			result = append(result, index, i)
			break
		}
	}
	return result
}
