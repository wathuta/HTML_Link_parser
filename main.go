package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/wathuta/HTML_Link_parser/parser"
)

func main() {
	port := flag.Int("port", 9090, "the port on which the server fires on")
	hh := parser.NewHTML(testHTML)

	mux := http.NewServeMux()
	mux.Handle("/", hh)
	log.Fatal(http.ListenAndServe("localhost:"+fmt.Sprintf("%d", *port), mux))

}

var testHTML = `<!DOCTYPE html>
<html lang="en">
<head>
	 <title>Document</title>
</head>
<body>
    <a href="www.google.com">GOOGLE</a>
    <a href="www.github.com">GITHUB</a>
    <a href="www.gitlab.com">GITLAB</a>
    <a href="www.youtube.com"><span>youtube</span></a>
</body>
</html>`
