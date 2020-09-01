package msg

import (
	"strings"
	"testing"
	"time"
)

func TestHelloMessage_MarshallMessage(t *testing.T) {
	tests := []struct {
		name    string
		hm      HelloMessage
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			string("Simple Hello World Test!"),
			HelloMessage{
				Msg:   "Hello World",
				Stamp: time.Now(),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.hm.MarshallMessage()
			if (err != nil) != tt.wantErr {
				t.Errorf("HelloMessage.MarshallMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if strings.Contains(string(got), tt.hm.Msg) != true {
				t.Errorf("HelloMessage.MarshallMessage() = %v, want %v", got, tt.hm.Msg)
			}
		})
	}
}
