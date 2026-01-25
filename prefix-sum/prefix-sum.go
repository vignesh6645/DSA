package prefixsum

import "fmt"

func PrefixSum() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	res1 := obj.SumRange(0, 2) // Expected: 1
	res2 := obj.SumRange(2, 5) // Expected: -1
	res3 := obj.SumRange(0, 5) // Expected: -3

	// Print results to verify
	fmt.Println(res1)
	fmt.Println(res2)
	fmt.Println(res3)
	fmt.Println("productExceptSelf", productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

type NumArray struct {
	prefixSums []int
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	sums := make([]int, n+1)

	for i := 0; i < n; i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	return NumArray{prefixSums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefixSums[right+1] - this.prefixSums[left]
}

func productExceptSelf(nums []int) []int {
	//Approach-1: Brute force solution
	// output := make([]int, len(nums))
	// for i := range nums {
	// 	out := 1
	// 	for j, n := range nums {
	// 		if i != j {
	// 			out *= n
	// 		}
	// 	}
	// 	output[i] = out
	// }
	// return output

	//Approach-2: Brute force solution

	n := len(nums)
	resultArr := make([]int, n)

	resultArr[0] = 1

	for i := 1; i < n; i++ {
		resultArr[i] = resultArr[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		resultArr[i] = resultArr[i] * right
		right = right * nums[i]
	}

	return resultArr

}
