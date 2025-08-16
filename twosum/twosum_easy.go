package twosum

func TwoSum_easy1() []int {
	nums := []int{2, 7, 11, 15}
	target := 9

	numsMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := numsMap[diff]; ok {
			return []int{j, i}
		}
		numsMap[num] = i
	}
	return []int{}
}
