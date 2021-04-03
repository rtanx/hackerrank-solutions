package mathematics

import (
	"fmt"
	"math/big"
)

func powerCalculation(k int64, n int64) string {
	sum, tmp := big.NewInt(int64(0)), big.NewInt(int64(0))
	var i int64
	for i = 1; i <= k; i++ {
		tmp.Exp(big.NewInt(i), big.NewInt(n), nil)
		sum.Add(sum, tmp)
	}
	fmt.Println(sum)
	// r := []rune(fmt.Sprintf("%02d", res))
	return "done"
}
