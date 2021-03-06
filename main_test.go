package main

// go get bitbucket.org/tebeka/go2xunit
// go test -v | $GOPATH/bin/go2xunit

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// statusHandler is an http.Handler that writes an empty response using itself
// as the response status code.
type statusHandler int

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(int(*h))
}

func TestIsTagged(t *testing.T) {
	// Set up a fake "Google Code" web server reporting 404 not found.
	status := statusHandler(http.StatusNotFound)
	s := httptest.NewServer(&status)
	defer s.Close()

	if isTagged(s.URL) {
		t.Fatal("isTagged == true, want false")
	}

	// Change fake server status to 200 OK and try again.
	status = http.StatusOK

	if !isTagged(s.URL) {
		t.Fatal("isTagged == false, want true")
	}

	time.Sleep(time.Second)
}

func TestIntegration(t *testing.T) {
	s := NewServer()

	// Make first request to the server.
	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	if b := w.Body.String(); !strings.Contains(b, "No.") {
		t.Fatalf("body = %s, want no", b)
	}
}
