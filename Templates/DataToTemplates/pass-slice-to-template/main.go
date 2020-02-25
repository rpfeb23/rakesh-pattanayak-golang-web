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

	inputstrings := []string{"A", "B", "C", "D"}

	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",inputstrings)

	if err != nil {
		log.Fatal(err)
	}


}