package server

import "sync"

// Server is a server
type Server struct {
	db *sync.Map
}
