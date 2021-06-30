package urlshort

import (
	"net/http"
)

/*
	Devuelve un http.HandlerFunc que mapea cualquier dirección con
	su URL
*/

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			Si tiene una dirección válida, nos vamos a ella
			si no, usamos el fallback
		*/
		path := r.URL.Path

		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}
