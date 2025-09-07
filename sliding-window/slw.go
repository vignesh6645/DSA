package slw

import "log"

func SlwInit() {
	log.Println("findMaxAverage: ", findMaxAverage([]int{0, 4, 0, 3, 2}, 1))
	log.Println("maxProfit: ", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	log.Println("CharacterReplacement", characterReplacement("ABAB", 2))
}

func characterReplacement(s string, k int) int {
	return k
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
