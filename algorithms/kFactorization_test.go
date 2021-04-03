package algorithms

import (
	"fmt"
	"testing"
)

func TestKFactorization(t *testing.T) {
	fmt.Println(kFactorization(663000000, []int32{2, 3, 5, 7, 11, 13, 17, 19}))
	// fmt.Println(kFactorization(12, []int32{3, 2, 4}))
	// fmt.Println(kFactorization(15, []int32{2, 10, 6, 9, 11}))
}
