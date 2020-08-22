package msg

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// ServeHello Serves Hello Message
func ServeHello() {
	log.Println(" DBG: Accepting traffic for /hello")
	helloHandler := http.HandlerFunc(HelloHandler)
	http.Handle("/hello", helloHandler)
}

//HelloHandler Handles the hello
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	var msg = "Hello World!"

	// Some simple debig messages
	log.Println(" DBG: Method - ", r.Method)
	log.Println(" DBG: Body - ", r.Body)
	log.Println(" DBG: URL - ", r.URL)
	m, err := url.ParseQuery(r.URL.RawQuery)

	if err == nil && len(m) != 0 {
		mm, ok := m["msg"]
		if ok != false {
			msg = mm[0]
			log.Println(" DBG: MSG - ", msg)
		}
	}

	// Stage 2 - Need to use the input to parse the message later
	hm := HelloMessage{
		Msg:   msg,
		Stamp: time.Now(),
	}

	sm, err := hm.MarshallMessage()
	if err != nil {
		log.Fatalf(" Error: %v", err)
		return
	}

	log.Printf(" DBG: %s", string(sm))
	io.WriteString(w, string(sm))
}
