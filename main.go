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
	listen := flag.String("listen", ":8080", "listen address")
	dir := flag.String("dir", ".", "directory to serve")
	flag.Parse()

	log.Printf("listening on %q...", *listen)
	log.Fatalln(http.ListenAndServe(*listen, logRequest(http.FileServer(http.Dir(*dir)))))
}
