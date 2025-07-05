package twosum

func TwoSum_easy1() [2]int {
	nums := []int{3, 2, 4}
	target := 6

	numsMap := make(map[int]int)

	result := [2]int{}

	for i, num := range nums {
		numsMap[num] = i
	}

	for i, num := range nums {
		diff := target - num
		if val, ok := numsMap[diff]; ok && val != i {
			result[0] = i
			result[1] = val
			break
		}
	}

	return result
}
