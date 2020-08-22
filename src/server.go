// server entry point
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/samarthya/msg"
)

func main() {
	log.Printf("%s", "My Simple Http Server")
	fmt.Println(" -------- Welcome to my server ------- ")
	msg.ServeHello()
	log.Fatal(http.ListenAndServe(":8090", nil))
}
