package redirect

import (
	"fmt"
	"net/http"
)

func Serve(location, port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", location+r.URL.Path)
		w.WriteHeader(301)
	})
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
