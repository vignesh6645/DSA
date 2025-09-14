package slw

import (
	"fmt"
	"log"
	"math"
)

func SlwInit() {
	log.Println("findMaxAverage: ", findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
	log.Println("maxProfit: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Println("CharacterReplacement: ", characterReplacement("BAAA", 0))
	log.Println("LengthOfLongestSubstring: ", lengthOfLongestSubstring("abcabcbb"))
	log.Println("findAnagrams: ", findAnagrams("cbaebabacd", "abc"))
	log.Println("MinWindow: ", minWindow("ADOBECODEBANC", "ABC"))
	log.Println("LongestSubstring: ", longestSubstring("ababbc", 2))
	log.Println("checkInclusion: ", checkInclusion("ab", "eidbaooo"))
	log.Println("maximumSubarraySum: ", maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
}

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	if k > n {
		return 0
	}

	maxVal := 100000
	freq := make([]int, maxVal+1)
	var currentSum int64 = 0
	var maxSum int64 = 0
	distinctCount := 0

	left := 0
	for right := 0; right < n; right++ {
		// Add nums[right]
		if freq[nums[right]] == 0 {
			distinctCount++
		}
		freq[nums[right]]++
		currentSum += int64(nums[right])

		// Shrink if window is bigger than k
		if right-left+1 > k {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinctCount--
			}
			currentSum -= int64(nums[left])
			left++
		}

		// Check if valid window
		if right-left+1 == k && distinctCount == k {
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}

	return maxSum
}

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freqS1 := make([]int, 26)     // frequency of chars in s1
	freqWindow := make([]int, 26) // frequency of current window in s2

	// Build frequency for s1
	for ind, ch := range s1 {
		freqS1[ch-'a']++
		freqWindow[s2[ind]-'a']++
	}

	windowSize := len(s1)

	// Compare the first window
	if equal(freqS1, freqWindow) {
		return true
	}

	// Slide the window
	for i := windowSize; i < len(s2); i++ {
		freqWindow[s2[i]-'a']++            // add new char
		freqWindow[s2[i-windowSize]-'a']-- // remove old char

		if equal(freqS1, freqWindow) {
			return true
		}
	}

	return false

}

func longestSubstring(s string, k int) int {
	maxLen := 0
	n := len(s)

	// Iterate over all possible number of unique characters (1 to 26)
	for h := 1; h <= 26; h++ {
		count := make([]int, 26)
		left, right := 0, 0
		unique, countAtLeastK := 0, 0

		for right < n {
			// Expand the window
			idx := s[right] - 'a'
			strIdx := string(s[right])
			fmt.Print(strIdx)
			if count[idx] == 0 {
				unique++
			}
			count[idx]++
			if count[idx] == k {
				countAtLeastK++
			}
			right++

			// Shrink the window if unique > h
			for unique > h {
				idxL := s[left] - 'a'
				if count[idxL] == k {
					countAtLeastK--
				}
				count[idxL]--
				if count[idxL] == 0 {
					unique--
				}
				left++
			}

			// Check if current window is valid
			if unique == h && unique == countAtLeastK {
				if right-left > maxLen {
					maxLen = right - left
				}
			}
		}
	}

	return maxLen
}

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	have := make(map[byte]int)
	required := len(need)
	formed := 0

	left := 0
	minLen := math.MaxInt32
	start := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		have[ch]++

		if need[ch] > 0 && have[ch] == need[ch] {
			formed++
		}

		// Try to shrink from the left
		for formed == required {
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			leftChar := s[left]
			have[leftChar]--
			if need[leftChar] > 0 && have[leftChar] < need[leftChar] {
				formed--
			}
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
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
	lastSeen := make([]int, 128)
	for i := range lastSeen {
		lastSeen[i] = -1
	}

	maxLength, start := 0, 0
	for i, ch := range s {
		if lastSeen[ch] >= start {
			start = lastSeen[ch] + 1
		}
		lastSeen[ch] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
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
