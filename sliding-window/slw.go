package slw

import (
	"fmt"
	"log"
)

func SlwInit() {
	log.Println("findMaxAverage: ", findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
	log.Println("maxProfit: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Println("CharacterReplacement: ", characterReplacement("BAAA", 0))
	log.Println("LengthOfLongestSubstring: ", lengthOfLongestSubstring("abcabcbb"))
	log.Println("findAnagrams: ", findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	result := []int{}
	pCount := make([]int, 26) // frequency of chars in p
	sCount := make([]int, 26) // frequency of chars in current window of s

	// Build frequency for p and first window in s
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	// Compare first window
	if equal(pCount, sCount) {
		result = append(result, 0)
	}

	// Slide the window
	for i := len(p); i < len(s); i++ {
		// add new char
		sCount[s[i]-'a']++
		// remove old char
		sCount[s[i-len(p)]-'a']--

		if equal(pCount, sCount) {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}

// helper: compare two frequency arrays
func equal(a, b []int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength, start := 0, 0

	for i, char := range s {
		strI := string(char)
		fmt.Print(strI)
		if lastSeenIndex, found := charIndexMap[char]; found && lastSeenIndex >= start {
			// Move the start pointer to the next position of the last seen duplicate character
			start = lastSeenIndex + 1
		}
		charIndexMap[char] = i
		// Calculate the current substring length
		currentLength := i - start + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}
	return maxLength
}

func characterReplacement(s string, k int) int {
	count := make([]int, 26) // instead of map
	left, right, maxCount, maxLength := 0, 0, 0, 0

	for right < len(s) {
		idx := s[right] - 'A'
		count[idx]++
		if count[idx] > maxCount {
			maxCount = count[idx]
		}

		// If window is invalid, shrink it
		if (right-left+1)-maxCount > k {
			count[s[left]-'A']--
			left++
		}

		// Update max length
		if (right - left + 1) > maxLength {
			maxLength = right - left + 1
		}

		right++
	}

	return maxLength
}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		currentProfit := prices[i] - minPrice
		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}
