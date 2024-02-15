package main

import (
	"log"
	"net/http"
	"os"
	"workspace/handlers"
)

func main() {

	l := log.New(os.Stdout, "Product API", log.LstdFlags)
	hh := handlers.NewHello(l)

	http.NewServeMux().Handle("/", hh)

	http.ListenAndServe(": 9090", nil)

}
