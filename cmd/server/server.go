package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/todoer-org/todoer/src/uuid"
)

const port = ":8080"

func main() {

	http.Handle("/", NewServer())
	log.Fatal(http.ListenAndServe(port, nil))

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

	uuid := uuid.NewUUID()

	fmt.Fprintf(w, "Hello, world! You have called me %d times %s.\n", s.hitCount, uuid)
}
