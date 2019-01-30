package main

import (
	"html/template"
	"os"
	"log"
)

var tpl *template.Template



// Keep in mind that we named the variable Name exported, if we named it  "name" we wouldn't have access to it in our
//template
type myStruct struct{
	Name string
	Last_name  string
	Favourtie_years []int
	Results map[string]int
}

func init(){
	tpl = template.Must(template.ParseGlob("template files/*.html"))
}



func main(){

	pooria := myStruct{
		"pooria",
		"poorsarvi",
		[]int{1, 2, 3},
		map[string]int{
			"math" : 20,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "main-template.html", pooria)
	if err != nil{
		log.Fatal(err)
	}
}



