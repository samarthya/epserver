// server entry point
package main

/**
 * Look for more contents [here](https://golang.org/doc/articles/wiki/)
 */
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/samarthya/msg"
	"github.com/samarthya/wiki"
)

var systemPort string

// Reference - https://golang.org/doc/effective_go.html#init
func init() {
	msg.ServeHello()
	msg.HealthServer()
	wiki.SetupView()

	systemPort = os.Getenv("SERVERPORT")

	if _, ok := strconv.ParseInt(systemPort, 10, 64); ok != nil {
		log.Println(" DBG - Setting Port (8090)")
		systemPort = "8090"
	}

	systemPort = fmt.Sprintf(":%s", systemPort)
	log.Println(" DBG: Port ", systemPort)
}

func main() {

	log.Printf("%s", "My Simple Http Server")
	fmt.Println(" -------- Welcome to my server ------- ")
	log.Fatal(http.ListenAndServe(systemPort, nil))
}
