package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
	"time"
)

type countHandler struct {
	mu sync.Mutex
	n  map[string]interface{}
}

var c = countHandler{n: map[string]interface{}{"counter": 0}}

func (c *countHandler) serveHTML(w http.ResponseWriter, filename string) {
	c.mu.Lock()
	c.n["counter"] = c.n["counter"].(int) + 1
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer c.mu.Unlock()
	if err := t.Execute(w, c.n); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	c.serveHTML(w, "static/index.html")
}

func main() {

	server := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(mainHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())

}
