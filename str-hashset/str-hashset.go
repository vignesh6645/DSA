package strhashset

import (
	"log"
	"strings"
	"unicode"
)

func Strhashset() {
	log.Println("containsDuplicate: ", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
	log.Println("anagram(): ", anagram("conversation", "voices rant on"))
	log.Println("groupAnagrams: ", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	log.Println("TopKFrequent: ", topKFrequent([]int{-1, -1}, 1))
	log.Println("IsSubsequence: ", isSubsequence("abc", "ahbgdc"))
	log.Println("LongestConsecutive: ", longestConsecutive([]int{1, 0, 1, 2}))
	log.Println("productExceptSelf: ", productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	left, right := make([]int, len(nums)), make([]int, len(nums))

	left[0] = 1
	right[len(nums)-1] = 1
	for i := 0; i < len(nums)-1; i++ {
		left[i+1] = nums[i] * left[i]
	}

	for i := len(nums) - 1; i > 0; i-- {
		right[i-1] = nums[i] * right[i]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = left[i] * right[i]
	}

	return nums
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]struct{})
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	maxLength := 0

	for num := range numSet {
		// Only expand if it's the start of a sequence
		if _, exists := numSet[num-1]; !exists {
			currentNum := num
			currentLength := 1

			for {
				_, exists := numSet[currentNum+1]
				if !exists {
					break
				}
				currentNum++
				currentLength++
			}

			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

func isSubsequence(s string, t string) bool {
	left, right := 0, 0
	for left < len(s) && right < len(t) {
		if s[left] == t[right] {
			left++
			right++
		} else {
			right++
		}
	}
	return left == len(s)
}

func topKFrequent(nums []int, k int) []int {
	lengthOfNums := len(nums)
	if lengthOfNums <= 1 {
		return nums
	}

	counterMap := make(map[int]int)
	for _, num := range nums {
		counterMap[num]++
	}

	bucket := make([][]int, lengthOfNums+1)
	for index, uniqueCounter := range counterMap {
		bucket[uniqueCounter] = append(bucket[uniqueCounter], index)
	}

	result := []int{}
	for i := lengthOfNums; i >= 0 && len(result) < k; i-- {
		if len(bucket[i]) > 0 {
			result = append(result, bucket[i]...)
		}
	}

	return result
}

func sortString(s string) string {
	count := [26]int{}
	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
	}

	var b []byte
	for i := 0; i < 26; i++ {
		for count[i] > 0 {
			b = append(b, byte('a'+i))
			count[i]--
		}
	}

	return string(b)
}

func groupAnagrams(strs []string) [][]string {
	groupOfSortedString := make(map[string][]string)

	for _, s := range strs {
		sortedString := sortString(s)
		groupOfSortedString[sortedString] = append(groupOfSortedString[sortedString], s)
	}

	result := make([][]string, 0, len(strs))
	for _, group := range groupOfSortedString {
		result = append(result, group)
	}
	return result
}

func containsDuplicate(nums []int) bool {

	if len(nums) == 0 {
		return false
	}

	resultMap := make(map[int]struct{})

	for _, num := range nums {

		if _, exist := resultMap[num]; exist {
			return true
		}
		resultMap[num] = struct{}{}

	}

	return false
}
func normalize(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

func anagram(s, t string) bool {
	s = normalize(s)
	t = normalize(t)

	if len(s) != len(t) {
		return false
	}

	bucket := make(map[rune]int)
	for _, ch := range s {
		bucket[ch]++
	}

	for _, ch := range t {
		if _, ok := bucket[ch]; !ok {
			return false
		}
		bucket[ch]--
		if bucket[ch] < 0 {
			return false
		}
	}

	return true
}
