package main

import (
	"log"
	"net/http"

	"srv/mdHandler"
    "srv/cfgParser"
)

func requestLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("%s\t%s", r.Method, r.URL.Path)
	})
}

func main() {
    cfgFields := cfgParser.Parse()
	mdHandler.HandleDir("data", cfgFields.Template)
	fileServer := http.FileServer(http.Dir("./site"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer) // Handle serves a handler at the given path
    log.Println("Starting server at port: " + cfgFields.Port)
	log.Fatal(http.ListenAndServe(":" + cfgFields.Port, requestLog(mux)))
}
