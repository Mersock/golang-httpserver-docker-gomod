package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	http.Handle("/", loggingMiddleware(http.HandlerFunc(hanler)))
	http.ListenAndServe(":8080", nil)
}

func hanler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("uir: %s", req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
