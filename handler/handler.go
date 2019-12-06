package handler

import (
	"net/http"
)
//this file is present in handler package which is imported by our main package which has our main.go file"
//exported function for handling the web requests of key value type
func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if _, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, path, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
