package msg

import (
	"io"
	"log"
	"net/http"
	"time"
)

var hm = HelloMessage{
	Msg:   HEALTHY,
	Stamp: time.Now(),
}

// HealthServer Serves Health Message
func HealthServer() {
	log.Println(" DBG: Accepting traffic for /hello")
	healthHandler := http.HandlerFunc(HealthHandler)
	http.Handle("/health", healthHandler)
}

// DBG: Method
func dumpMessages(r *http.Request) {
	// Some simple debug messages
	log.Println(" DBG: Method - ", r.Method)
	log.Println(" DBG: Body - ", r.Body)
	log.Println(" DBG: URL - ", r.URL)
}

//HealthHandler Handles the hello
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// Dump debug messages
	dumpMessages(r)

	hm.Stamp = time.Now()

	sm, err := hm.MarshallMessage()
	if err != nil {
		log.Fatalf(" Error: %v", err)
		return
	}

	log.Printf(" DBG: %s", string(sm))
	io.WriteString(w, string(sm))
}
