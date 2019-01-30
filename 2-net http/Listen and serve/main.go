package main

import (
	"net/http"
	"io"
)

type anything int

func (m anything) ServeHTTP(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "how you doiin !")
}

func main(){

	var d anything

	http.ListenAndServe(":8080", d)

}


