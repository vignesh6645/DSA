package twosum

var (
	nums   = []int{1, 2, 3, 4, 6, 7}
	target = 9
)

func TwoSum_easy1() []int {

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

func TwoSum_SortedArr() bool {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
