package main

import (
	"sync"

	"github.com/ianchen06/airflow-backfiller/server"
)

type cool struct {
	db *sync.Map
}

func main() {
	s := server.Server{db: &sync.Map{}}
	c := cool{db: &sync.Map{}}
}
