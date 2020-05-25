package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct
type Goodbye struct {
	logger *log.Logger
}

// NewGoodbye constructor
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Byeee"))
}
