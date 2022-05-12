package main

import (
	"flag"
	"log"
	"net/http"
)

type Config struct {
	Addr      string
	StaticDir string
}

func main() {
	config := new(Config)
	flag.StringVar(&config.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&config.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix(("/static"), fileServer))

	log.Printf("Starting server on %s", config.Addr)
	err := http.ListenAndServe(config.Addr, mux)
	log.Fatal(err)
}
