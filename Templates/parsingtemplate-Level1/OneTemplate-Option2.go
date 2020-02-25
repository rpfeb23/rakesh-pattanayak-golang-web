// Single Template File using
//     1. template.ParseFiles (template is Package) retruns pointer to Template use that to Execute
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

	newfile, err := os.Create("index-option2.html")
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()

	err = tpl.Execute(newfile,tpl)
	if err != nil {
		log.Fatal(err)
	}

}