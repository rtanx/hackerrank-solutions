package algorithms

import (
	"fmt"
	"sort"
)

func kFactorization(n int32, A []int32) []int32 {
	sort.SliceStable(A, func(i, j int) bool {
		return A[i] > A[j]
	})
	dTmp := n
	var fctr []int32
	// fctr := []int32{1, n}
	for _, v := range A {
		if dTmp%v != 0 {
			continue
		}
		dTmp = dTmp / v
		fmt.Println(dTmp)
		fctr = append(fctr, dTmp)
	}
	if len(fctr) == 0 {
		return []int32{-1}
	}
	fctr = append(fctr, n)
	sort.SliceStable(fctr, func(i, j int) bool {
		return fctr[i] < fctr[j]
	})
	return fctr

}

func factor(fctr []int32) int32 {
	res := int32(1)
	for _, f := range fctr {
		res *= f
	}
	return res
}
