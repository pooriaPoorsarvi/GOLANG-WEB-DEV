package main

import (
	"html/template"
	"net/http"
	"log"
	"net/url"
)

var tpl *template.Template


func init(){

	tpl = template.Must(template.ParseGlob("template files/*.html"))


}


type intermid int

func (m intermid) ServeHTTP(w http.ResponseWriter, req *http.Request){
	err := req.ParseForm()

	if err != nil{
		log.Println(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Submission map[string][]string
		Header http.Header
		Host string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}


	tpl.ExecuteTemplate(w, "request-info.html", data)


}

func main(){

	var in1 intermid

	http.ListenAndServe(":8080", in1)



}






