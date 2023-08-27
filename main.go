package main

import (
	"log"
	"net/http"
	"path"
	"strings"

	"srv/cfgParser"
	"srv/mdHandler"
)

func requestLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Printf("%s\t%s", r.Method, r.URL.Path)
	})
}

func renderHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func renderBlog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "blog.html")
}

func serveBlogPage(w http.ResponseWriter, r *http.Request) {
	// Extract folderName from URL
	parts := strings.Split(r.URL.Path, "/")
	log.Println("hereh", parts, ",len,", len(parts))
	if len(parts) < 3 {
		http.NotFound(w, r)
		return
	}
	folderName := parts[2]
	log.Println("hereh", parts, ",len,", len(parts), ",fname,", folderName)

	// Serve requested page
	pagePath := path.Join("./site/", folderName, "index.html")
	http.ServeFile(w, r, pagePath)
}

func main() {
	cfgFields := cfgParser.Parse()

	mdHandler.HandleDir("data", cfgFields.Template, "blog_layout.html")

	//fileServer := http.FileServer(http.Dir("./site"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", renderHome)         // Handle serves a handler at the given path
	mux.HandleFunc("/blog", renderBlog)     // Handle serves a handler at the given path
	mux.HandleFunc("/blog/", serveBlogPage) // Handle serves a handler at the given path
	log.Println("Starting server at port: " + cfgFields.Port)
	log.Fatal(http.ListenAndServe(":"+cfgFields.Port, requestLog(mux)))
}
