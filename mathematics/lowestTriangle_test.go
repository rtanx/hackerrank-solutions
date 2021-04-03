package mathematics

import (
	"testing"
)

func TestLowestTriangle(t *testing.T) {
	var assertions = []struct {
		b, a, h int
	}{
		{17, 100, 12},
		{2, 2, 2},
		{4, 6, 3},
	}
	for _, assert := range assertions {
		if h := LowestTriangle(assert.b, assert.a); h != assert.h {
			t.Errorf("The lowest height for a triangle with a base: %v and area: %v should be %v, but %v is returned", assert.b, assert.a, assert.h, h)
		}
	}
}
