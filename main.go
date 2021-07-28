package main

import (
	"flag"
	"log"
	"net/http"
)

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s: %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	l := flag.String("l", ":8080", "listen address")
	d := flag.String("d", ".", "directory to serve")
	flag.Parse()

	log.Printf("listening on %q...", *l)
	log.Fatal(http.ListenAndServe(*l, logRequest(http.FileServer(http.Dir(*d)))))
}
