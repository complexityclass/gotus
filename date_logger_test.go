package gotus

import "testing"

func TestLogCurrentTime(t *testing.T) {
	if got := LogCurrentTime(); got == UIError {
		t.Errorf("got: %s", got)
	}
}

func TestHelloNow(t *testing.T) {
	HelloNow()
}
