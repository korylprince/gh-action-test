package handler

import (
	"fmt"
	"net/http"
	"regexp"
)

var nameRegexp = regexp.MustCompile("^/([a-zA-z]+)$")

func HelloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		names := nameRegexp.FindStringSubmatch(r.URL.Path)
		if len(names) != 2 || names[1] == "" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", names[1])))
	})
}
