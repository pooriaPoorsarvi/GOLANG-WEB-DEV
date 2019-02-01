package main

import (
	"net/http"
	"io"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "hoooot daawwg")
}

type hotcat int

func (m hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "hoooot caawwt")
}




func main(){

	mux := http.NewServeMux()

	var d hotdog
	var c hotcat

	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)

}



