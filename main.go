package main

import (
	"log"

	strhashset "github.com/vignesh6645/dsa/str-hashset"
	"github.com/vignesh6645/dsa/twosum"
)

func main() {
	log.Println("Two_Sum_easy1: ", twosum.TwoSum_easy1())
	log.Println("TwoSum_SortedArr: ", twosum.TwoSum_SortedArr())
	log.Println("MoveZeroes: ", twosum.MoveZeroes())
	log.Println("MajorityElement: ", twosum.MajorityElement())
	log.Println("TwoSum: ", twosum.RemoveDuplicates())
	log.Println("Merge Sorted Array: ", twosum.Merge())
	strhashset.Strhashset()
}
