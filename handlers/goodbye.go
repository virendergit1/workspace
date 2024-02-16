package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

// Idiomatic. this is like constructor which works on a Struct defined above
func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

// now the methods
func (h *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "Hello from GB hadler %s\n", data)
}
