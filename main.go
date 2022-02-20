package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	indexIfNotFound := os.Getenv("INDEX_IF_NOT_FOUND")

	var fs http.Handler
	if indexIfNotFound != "" {
		fs = fileServer(http.Dir("/app"))
	} else {
		fs = http.FileServer(http.Dir("/app"))
	}

	http.Handle("/", fs)

	log.Println("Listening on :80")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func fileServer(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			// Set the url path to / to fallback to index if no file is found
			r.URL.Path = "/"
		}
		fsh.ServeHTTP(w, r)
	})
}
