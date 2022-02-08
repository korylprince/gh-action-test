package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/korylprince/gh-action-test/handler"
)

func TestHelloHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/Kory", nil)

	rr := httptest.NewRecorder()
	h := handler.HelloHandler()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Check OK status: want: %d, have: %d", http.StatusOK, rr.Code)
	}

	if rr.Body.String() != "Hello, Kory!" {
		t.Errorf(`Check correct name: want: "%s", have: "%s"`, "Hello, Kory!", rr.Body.String())
	}

	req, _ = http.NewRequest("GET", "/bad/name", nil)

	rr = httptest.NewRecorder()
	h = handler.HelloHandler()
	h.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Check Not Found status: want: %d, have: %d", http.StatusNotFound, rr.Code)
	}
}
