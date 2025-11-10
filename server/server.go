package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"net/http"
)

// todo render markdown files in advance and serve static html files (?)
//
//go:embed pages/index.md
var rootPage string

//go:embed template.html
var templateHTML string

// to be adapted as a generic page rendering function
func root(w http.ResponseWriter, req *http.Request) {
	// ref: https://github.com/yuin/goldmark
	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithUnsafe(),
		),
	)
	var buf bytes.Buffer
	if err := md.Convert([]byte(rootPage), &buf); err != nil {
		http.Error(w, "Failed to render markdown", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	renderedHtml := fmt.Sprintf(templateHTML, buf.String())

	_, err := w.Write([]byte(renderedHtml))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	//w.Write([]byte("Hello, World!\n"))
	fmt.Println("/hello called")

	ctx := req.Context()
	select {
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("Request cancelled:", err)
	default:
		_, err := w.Write([]byte("sup"))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}
}

func StartServer() {
	fmt.Println("Starting server on :6969")

	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)

	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":6969", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	StartServer()
}
