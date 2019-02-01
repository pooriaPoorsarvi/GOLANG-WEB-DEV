package main

import (
	"net/http"
	"io"
)


func hotdog(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "hoooot daawwg")
}


func hotcat(w http.ResponseWriter, req *http.Request){
	io.WriteString(w, "hoooot caawwt")
}




func main(){


	//seriously check how good and simple the implementation is, this is the beauty of golang
	http.Handle("/dog", http.HandlerFunc(hotdog))
	http.Handle("/cat", http.HandlerFunc(hotcat))

	http.ListenAndServe(":8080", nil)

}



