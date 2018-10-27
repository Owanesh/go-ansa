package main

import (
	"log"
	"net/http"

	"github.com/Owanesh/go-ansa/pkg/fetching"
	"github.com/Owanesh/go-ansa/pkg/http/rest"
	"github.com/Owanesh/go-ansa/pkg/storage"
)

func main() {
	s := new(storage.Storage)
	service := fetching.NewService(s)

	router := rest.Handler(service)
	log.Fatal(http.ListenAndServe(":8080", router))
}
