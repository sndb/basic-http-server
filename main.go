package main

import (
	"flag"
	"log"
	"net/http"
)

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	var addr string
	flag.StringVar(&addr, "a", ":8080", "address")
	flag.Parse()
	dir := flag.Arg(0)
	if dir == "" {
		dir = "."
	}
	srv := http.FileServer(http.Dir(dir))
	log.Printf("Serving %s on %s", dir, addr)
	log.Fatal(http.ListenAndServe(addr, logRequest(srv)))
}
