package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io"
	"fmt"
)

func defaultHandler(w http.ResponseWriter, req *http.Request)  {
	io.WriteString(w, "this is the default handler.")
}


func userHandler(w http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)
	fmt.Fprintln(w, "hey ", vars["name"], "its great to have a ", vars["age"], " old here")
}


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/user/{name}/{age:[0-9]+}/", userHandler)
	r.PathPrefix("/").HandlerFunc(defaultHandler)

	http.ListenAndServe(":8080", r)

}



