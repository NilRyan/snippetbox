package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
}
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	config := new(Config)
	flag.StringVar(&config.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&config.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     config.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", config.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
