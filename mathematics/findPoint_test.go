package mathematics

import (
	"reflect"
	"testing"
)

func TestFindPoint(t *testing.T) {
	var assertion = []struct {
		pPoint         []int32
		qPoint         []int32
		reflectedPoint []int32
	}{
		{pPoint: []int32{0, 0}, qPoint: []int32{1, 1}, reflectedPoint: []int32{2, 2}},
		{pPoint: []int32{1, 1}, qPoint: []int32{2, 2}, reflectedPoint: []int32{3, 3}},
		{pPoint: []int32{7, 4}, qPoint: []int32{5, 8}, reflectedPoint: []int32{3, 12}},
	}

	for _, assert := range assertion {
		out := FindPoint(assert.pPoint[0], assert.pPoint[1], assert.qPoint[0], assert.qPoint[1])
		if !reflect.DeepEqual(out, assert.reflectedPoint) {
			t.Fail()
		}
	}
}
