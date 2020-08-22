package msg

import (
	"testing"
	"time"
)

// TestMarshalling It tests the JSON marshalling
func TestMarshalling(t *testing.T) {
	h := HelloMessage{
		Msg:   "Hello World",
		Stamp: time.Now(),
	}

	mb, err := h.MarshallMessage()

	if err != nil {
		t.Errorf(" Failed Marshalling, unexpected %v", err)
	}

	t.Logf(" Test Passed %v", string(mb))
}
