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

	// Pass Nil
	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",nil)
	if err != nil {
		log.Fatal(err)
	}

	// Pass something

	err = tpl.ExecuteTemplate(os.Stdout,"template1.html","James Bond")

	if err != nil {
		log.Fatal(err)
	}


}