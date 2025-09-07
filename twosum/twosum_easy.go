package twosum

import (
	"log"
)

var (
	nums   = []int{2, 3, 4}
	target = 6
	nums1  = []int{1, 2, 3, 0, 0, 0}
	nums2  = []int{2, 5, 6}
)

func TwoSumDsa() {
	log.Println("Two_Sum_easy1: ", TwoSum_easy1())
	log.Println("TwoSum_SortedArr: ", TwoSum_SortedArr())
	log.Println("MoveZeroes: ", MoveZeroes())
	log.Println("MajorityElement: ", MajorityElement())
	log.Println("TwoSum: ", RemoveDuplicates())
	log.Println("Merge Sorted Array: ", Merge())
	log.Println("isPalindrome: ", isPalindrome(" "))
	//log.Println("ThreeSum: ", threeSum([]int{-1, 0, 1, 2, -1, -4}))
	log.Println("MaxArea: ", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))

}

func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1

	for left < right {
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}

		area := h * (right - left + 1) // careful: width must match old positions
		if area > res {
			res = area
		}
	}

	return res
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/* func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	tmp := 0
	left, right := 1, len(nums)-1
	result := [][]int{}
	for left != right {
		if nums[tmp]+nums[left]+nums[right] == 0 {
			result = append(result, []int{nums[tmp], nums[left], nums[right]})
			left++
			right = len(nums) - 1
		} else if left <= len(nums)-1 && left == right {
			tmp++
			left++
			right = len(nums) - 1
		} else {
			right--
		}
	}
	return result
} */

func isPalindrome(s string) bool {
	b := []byte(s)
	n := 0

	for i := 0; i < len(b); i++ {
		c := b[i]
		if (c >= '0' && c <= '9') ||
			(c >= 'a' && c <= 'z') ||
			(c >= 'A' && c <= 'Z') {
			if c >= 'A' && c <= 'Z' {
				c = c + ('a' - 'A')
			}
			b[n] = c
			n++
		}
	}

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if b[i] != b[j] {
			return false
		}
	}

	return true

}

func TwoSum_easy1() []int {

	numsMap := make(map[int]int)
	for i, num := range nums {
		diff := target - num
		if j, ok := numsMap[diff]; ok {
			return []int{j + 1, i + 1}
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

func MoveZeroes() []int {
	lengthOfNums := len(nums)
	if lengthOfNums <= 1 {
		return nums
	}
	left, right := 0, 0

	for right < lengthOfNums {
		if nums[right] == 0 && nums[left] == 0 {
			right++
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right++
		}
	}
	return nums
}

func MajorityElement() int {
	canditate, count := 0, 0

	for _, num := range nums {

		if count == 0 {
			canditate = num
		}

		if num == canditate {
			count++
		} else {
			count--
		}
	}
	return canditate
}

func RemoveDuplicates() int {

	lengthOfNums := len(nums)
	if lengthOfNums < 1 {
		return 0
	}

	if lengthOfNums == 1 {
		return 1
	}

	left := 0

	for right := 1; right < lengthOfNums; right++ {
		if nums[right] != nums[left] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}

func Merge() []int {
	m, n := 3, 3
	last := m + n - 1
	//Merge array in reverse order
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			log.Println("*** num1last: ******", nums1[last])
			log.Println("***  nums1[m-1]: *******", nums1[m-1])
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}
		last--
	}
	// Merge remaining elements from num2 arr
	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}
	return nums1
}
