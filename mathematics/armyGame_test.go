package mathematics

import (
	"fmt"
	"testing"
)

func TestGameWithCells(t *testing.T) {
	var assertions = []struct {
		n int32 // rows
		m int32 // columns
		s int32 // should return
	}{
		{2, 2, 1},
		{4, 6, 6},
		{4, 2, 2},
	}
	for _, assert := range assertions {
		if s := GameWithCells(assert.n, assert.m); s != assert.s {
			fmt.Println(s, assert.s)
			t.Fail()
		}
	}
}
