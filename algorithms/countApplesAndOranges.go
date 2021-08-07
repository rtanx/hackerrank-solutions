package algorithms

import (
	"fmt"
)

func CountApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var ncnt, mcnt int32
	for _, m := range apples {
		if (a+m) >= s && (a+m) <= t {
			mcnt++
		}
	}
	for _, n := range oranges {
		if (b+n) >= s && (b+n) <= t {
			ncnt++
		}
	}
	fmt.Println(mcnt)
	fmt.Println(ncnt)
}
