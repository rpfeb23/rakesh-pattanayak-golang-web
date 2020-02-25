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

	err := tpl.ExecuteTemplate(os.Stdout,"template1.html","Miss Moneypenny")

	if err != nil {
		log.Fatal(err)
	}


}