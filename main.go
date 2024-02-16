package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"workspace/handlers"
)

func main() {

	l := log.New(os.Stdout, "Product API ->", log.LstdFlags)

	// Handers here
	//hh := handlers.NewHello(l)
	//gbh := handlers.NewGoodBye(l)
	ph := handlers.NewProductsHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	//sm.Handle("/hello", hh)
	//sm.Handle("/goodbye", gbh)

	// HTTP Server
	server := &http.Server{
		Addr:         ": 9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// do not block
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// create a channel to listen the OS events
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, os.Interrupt)
	signal.Notify(sigChannel, os.Kill)

	// block here
	sig := <-sigChannel
	l.Println("Received terminate, graceful shutdown", sig)
	// and then shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(tc)

}
