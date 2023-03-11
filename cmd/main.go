package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sujeetchnp/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// fmt.Printf("Starting server at port 9000\n")            // mysql
	// log.Fatal(http.ListenAndServe("localhost:9000", r))	  // mysql
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
