package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const port = 8080

func main() {

	fmt.Printf("Started server at http://localhost%v.\n", port)
	http.Handle("/", NewServer())
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))

}

// NewServer constructs a server instance
func NewServer() *Server {

	server := &Server{
		version: "1",
	}

	return server
}

// Server represents a http server
type Server struct {
	version  string
	hitCount int64
	mu       sync.RWMutex
}

// ServeHTTP method recives all the http traffic and is the root http handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.RLock()
	s.hitCount = s.hitCount + 1
	s.mu.RUnlock()

	fmt.Fprintf(w, "Hello, world! You have called me %d times.\n", s.hitCount)
}
