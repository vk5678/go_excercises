package handler

import (
	"net/http"
)

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
