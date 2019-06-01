package mtimelog

import (
	"testing"
)

func TestLogCurrentTime(t *testing.T) {
	if _, err := LogCurrentTime(); err != nil {
		t.Errorf("got: %s", err)
	}
}
