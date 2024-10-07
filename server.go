package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	port := flag.String("p", "8000", "port to listen on")
	flag.Parse()

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	log.Printf("Listening on port %s", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
