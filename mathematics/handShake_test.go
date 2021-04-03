package mathematics

import (
	"testing"
)

func TestHandShacke(t *testing.T) {
	var assertions = []struct {
		p  int32
		hs int32
	}{
		{5, 10}, {7999, 31988001},
		{18, 153}, {1, 0}, {0, 0},
	}
	for _, assert := range assertions {
		if hs := HandShake(assert.p); hs != assert.hs {
			t.Errorf("for %v of people, should %v handshakes done, but got %v", assert.p, assert.hs, hs)
		}
	}
}
