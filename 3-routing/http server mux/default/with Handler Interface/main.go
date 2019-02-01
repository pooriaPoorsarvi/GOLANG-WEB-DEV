package with_Handler_Interface

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


	var d hotdog
	var c hotcat

	http.Handle("/dog", d)
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil)

}



