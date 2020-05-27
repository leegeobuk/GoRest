package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/leegeobuk/GoRest/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProduct(l)

	mux := mux.NewRouter()

	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProducts)
	postRouter.Use(ph.MiddlewareProductAuthentication)

	putRouter := mux.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductAuthentication)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, shutting down gracefully:", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
