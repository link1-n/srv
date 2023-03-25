package main

import (
	// "fmt"
	// "html/template"
	"log"
	"net/http"

	// "os"

	"srv/mdHandler"
	// "github.com/russross/blackfriday/v2"
)

// func root(w http.ResponseWriter, r *http.Request) {
// data, tmpl := mdHandler.HandleFile("test.md")
// tmpl.Execute(w, data)
// }

// func readme(w http.ResponseWriter, r *http.Request) {
	// data, tmpl := mdHandler.HandleFile("README.md")
	// tmpl.Execute(w, data)
// }

func requestLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("%s\t%s", r.Method, r.URL.Path)
	})
}

func main() {
	//mdHandler.HandleFile("test.md")
	mdHandler.HandleDir("data")
	// fileServer is a handler that makes that path of the server as file server.
	fileServer := http.FileServer(http.Dir("./site"))
	// ????? what does mux do???
	mux := http.NewServeMux()
	// mux.HandleFunc("/", root) // HandleFunc serves a function of input type ( http.ResponseWrite, *http.Request) at the given path
	//mux.HandleFunc("/readme", readme)
	// fsHandler := http.StripPrefix("/files", fileServer)
	//mux.Handle("/files", fsHandler) // Handle serves a handler at the given path
	mux.Handle("/", fileServer) // Handle serves a handler at the given path
	/*
		if we don't use fsHandler the FileServer will get requests for files in the form of /files/<file>
		this will not work as the directory is / and the actual path of the file is /<file>
		thus we use StripPrefix to remove the prefix "/files" from the path of the file
	*/
	log.Println("running")
	//log.Fatal(http.ListenAndServe(":8080", mux))
	log.Fatal(http.ListenAndServe(":8080", requestLog(mux)))
}
