package main

import (
	"net/http"
	"io"
)

type interMid int


func (m interMid) ServeHTTP(w http.ResponseWriter, req *http.Request){

	switch req.URL.Path {
	case "/dog" :
		io.WriteString(w, "u hoo doggy")
	case "/cat" :
		io.WriteString(w, "yo kitty")
	default:
		io.WriteString(w, "u shouldn't be here")
	}

}


func main(){
	var d interMid

	http.ListenAndServe(":8080", d)

}



