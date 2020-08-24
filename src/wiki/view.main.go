package wiki

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var regExp = regexp.MustCompile(`\w*\.html`)

// SetupView Sets up the default view HTML
func SetupView() {
	defHandler := http.HandlerFunc(defaultHandler)
	http.Handle("/", defHandler)
	log.Println(" DBG: Accepting for /any.html")
}

//defaultHandler Handles the hello
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	title := strings.ToLower(string(regExp.Find([]byte(r.URL.Path[1:]))))

	if title == "" {
		title = "welcome.html"
	}
	p, err := LoadPage(title)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "<h1>%s</h1>%s", strings.ToUpper(p.Title), p.Body)
}

// Reference - https://blog.golang.org/error-handling-and-go
