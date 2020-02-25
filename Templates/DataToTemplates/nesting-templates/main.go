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

	tpl, err := tpl.ParseFiles("template2.txt")

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout,"template1.html","Miss Moneypenny")

	if err != nil {
		log.Fatal(err)
	}


}