package apitests

import (
	"net/http"
	"testing"
)

const (
	port = ":8080"
	host = "todoer-server"
)

func checkStatus(t *testing.T, r *http.Response) {
	if r.StatusCode != http.StatusOK {
		t.Errorf("response was not a 200(OK) ")
	}
}

func TestHelloWorldRequest(t *testing.T) {
	url := "http://" + host + port
	resp, err := http.Get(url)

	if err != nil {
		t.Fatalf("Failed to get: %s %v", url, err)
	}

	checkStatus(t, resp)
}
