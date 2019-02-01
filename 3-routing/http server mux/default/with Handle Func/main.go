package with_Handler_Interface

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



	http.HandleFunc("/dog", hotdog)
	http.HandleFunc("/cat", hotcat)

	http.ListenAndServe(":8080", nil)

}



