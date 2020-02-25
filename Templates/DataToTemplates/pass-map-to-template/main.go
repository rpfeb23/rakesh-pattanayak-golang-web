package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseGlob("template1.html"))
}

func main()  {

	inputmap := map[string]int{
		"Rakesh"  : 35,
		"Rajesh"  : 37,
		"Rojalin" : 40,
	}

	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",inputmap)

	if err != nil {
		log.Fatal(err)
	}


}