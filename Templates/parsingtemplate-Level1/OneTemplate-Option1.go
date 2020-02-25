// Single Template File using
//     1. template.ParseFiles (template is Package)
//     2. Template.Execute (Template is a Type)
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main()  {
	tpl,err := template.ParseFiles("file1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(tpl)

	err = tpl.Execute(os.Stdout,tpl)
	if err != nil {
		log.Fatal(err)
	}

}