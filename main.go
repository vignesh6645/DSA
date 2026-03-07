package main

import (
	prefixsum "github.com/vignesh6645/dsa/prefix-sum"
	slw "github.com/vignesh6645/dsa/sliding-window"
	strhashset "github.com/vignesh6645/dsa/str-hashset"
	"github.com/vignesh6645/dsa/tree"
	"github.com/vignesh6645/dsa/twosum"
)

func main() {
	strhashset.Strhashset()
	twosum.TwoSumDsa()
	slw.SlwInit()
	prefixsum.PrefixSum()

	//var mapp map[string]string
	tree.InitBTree()

}
