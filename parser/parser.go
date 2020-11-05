package parser

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//Link is a model of the output
type Link struct {
	Href string
	Text string
}

//HTMLhandler is the handler that handles parsing of the html
type HTMLhandler struct {
	t string
}

//NewHTML is the entry point to the handler
func NewHTML(t string) http.Handler {
	return &HTMLhandler{t}
}
func (h *HTMLhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	docs, err := html.Parse(strings.NewReader(h.t))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					var link Link
					link.Href = attr.Val
					link.Text = n.Data
					w.Write([]byte(fmt.Sprintf("%v\t", link.Text)))
					w.Write([]byte(fmt.Sprintf("%v\n", link.Href)))
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(docs)

}
