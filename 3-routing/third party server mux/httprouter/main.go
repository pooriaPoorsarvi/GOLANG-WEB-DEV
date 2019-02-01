package httprouter

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
)

func index(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	fmt.Fprintln(w, "you are in index !!")
}

func user(w http.ResponseWriter, req *http.Request, params httprouter.Params){
	fmt.Fprintln(w, "Hello there", params.ByName("name"))
}

func main(){

	mux := httprouter.New()
	mux.GET("/", index)
	mux.GET("/user/:name", user)

	http.ListenAndServe(":8080", mux)



}



