package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

func main(){
	r := mux.NewRouter()

	r.PathPrefix("/resource").Handler(http.StripPrefix("/resource", http.FileServer(http.Dir("./assets"))))

	log.Fatal(	http.ListenAndServe(":8080", r) )

}


