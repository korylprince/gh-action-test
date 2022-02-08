package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/korylprince/gh-action-test/handler"
)

func main() {
	listen := os.Getenv("LISTENADDR")
	if listen == "" {
		listen = ":80"
	}

	http.Handle("/", handler.HelloHandler())
	if err := http.ListenAndServe(listen, nil); err != nil {
		fmt.Println("could not run server:", err)
		os.Exit(1)
	}
}
