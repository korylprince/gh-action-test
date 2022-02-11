package handler

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/google/uuid"
)

var nameRegexp = regexp.MustCompile("^/([a-zA-z]+)$")

// HelloHandler is a http.Handler that says "Hello, name!"
func HelloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		names := nameRegexp.FindStringSubmatch(r.URL.Path)
		if len(names) != 2 || names[1] == "" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", names[1])))
		log.Println("UUID:", uuid.NewString())
	})
}
