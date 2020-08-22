package msg

import (
	"encoding/json"
	"log"
	"time"
)

/**
 * Reference
 * - https://blog.golang.org/json
 * - https://golang.org/pkg/net/url/#Values
 * - https://golang.org/pkg/time/#pkg-examples
 */

// HelloMessage Hello Message struct
type HelloMessage struct {
	Msg   string    `json:"message"`
	Stamp time.Time `json:"timestamp"`
}

// MarshallMessage Marshals the message
func (hm HelloMessage) MarshallMessage() ([]byte, error) {

	log.Printf("\t{ ")
	log.Printf("\t DBG: %s", hm.Msg)
	log.Printf("\t DBG: %s", hm.Stamp.Local())
	log.Printf("\t} ")

	mh, err := json.Marshal(hm)
	if err == nil {
		return mh, nil
	}
	// Returns an empty byte
	return []byte(""), err
}
