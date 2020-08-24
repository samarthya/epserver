package wiki

import (
	"io/ioutil"
	"strings"
)

// Reference - https://golang.org/doc/articles/wiki/

// Page defines the structure.
type Page struct {
	Title string
	Body  []byte
}

// Save function
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// LoadPage - The function loadPage constructs the file name from the title parameter,
// reads the file's contents into a new variable body,
// and returns a pointer to a Page literal constructed
// with the proper title and body values.
func LoadPage(title string) (*Page, error) {
	filename := "./templates/" + title
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: strings.TrimRight(title, ".html"), Body: body}, nil
}
