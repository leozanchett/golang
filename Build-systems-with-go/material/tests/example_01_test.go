package examples

import "testing"

func TestMe(t *testing.T) {
	r := 2 + 2
	if r != 4 {
		t.Errorf("2+2 should be 4, but it is %d", r)
	}
}
