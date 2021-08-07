package algorithms

import (
	"sort"
)

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	sort.Slice(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	var pairs int32
	for i, x := range ar {
		for j := (int32(i) + 1); j < n; j++ {
			if (x+ar[j])%k == 0 {
				pairs++
			}
		}
	}
	return pairs
}
