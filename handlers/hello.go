package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// this is like constructor which works on a Struct defined above
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// now the methods
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello from h handler %s\n", data)
}
