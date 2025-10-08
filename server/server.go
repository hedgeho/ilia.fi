package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/yuin/goldmark"
	"net/http"
)

//go:embed pages/index.md
var rootPage string

func root(w http.ResponseWriter, req *http.Request) {
	md := goldmark.New()
	var buf bytes.Buffer
	if err := md.Convert([]byte(rootPage), &buf); err != nil {
		http.Error(w, "Failed to render markdown", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := w.Write(buf.Bytes())
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
	}
}

func StartServer() {
	http.HandleFunc("/", root)
	http.HandleFunc("/hello", hello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func main() {
	StartServer()
}
