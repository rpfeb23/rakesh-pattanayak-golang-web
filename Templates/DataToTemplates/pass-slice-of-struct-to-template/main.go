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

	inputstructslice := []struct{
		Name string
		Age int
	}{struct {
		Name string
		Age  int
		}{Name: "James Bond", Age:40 },
		struct {
			Name string
			Age  int
		}{Name: "Rakesh ", Age: 35 },
	}

	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",inputstructslice)

	if err != nil {
		log.Fatal(err)
	}


}