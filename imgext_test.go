package imgext

import "testing"

func TestGet(t *testing.T) {
	n := Get()
	if len(n) <= 0 {
		t.Fatalf("There must be atleast one extension provided")
	}
}
