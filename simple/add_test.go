package simple

import "testing"

func TestAdd(t *testing.T) {
	t.Logf("test Add func")
	c := Add(10, 11)
	if c == 21 {
		t.Logf("test Add 10+11 success")
	} else {
		t.Errorf("test Add 10+11 fail, expected 21, but given %d", c)
	}
	c = Add(1, 2)
	if c != 3 {
		t.Errorf("test Add 1+2 fail, expected 3, but given %d", c)
	}
}
