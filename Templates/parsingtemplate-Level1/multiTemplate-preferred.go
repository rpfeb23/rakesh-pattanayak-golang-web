// Single Template File using
//     1. Add templates using
//        template.ParseGlob (template is Package)
//        you can specifiy the file path
//        retruns pointer
//        to Template use that to Execute
//     2. Add additional templates using Template.ParseFiles
//     3. Template.Execute (Template is a Type) will randomly pick 1
//     4. Template.ExecuteFiles will execute Specific template file
package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseGlob("templates/*.text"))
	// if you will just do templates/* all files under templates folder will be picked
}

func main()  {
	// Ranomly executes any template even though tpl has reference to all 3 template files
	err := tpl.Execute(os.Stdout,tpl)
	if err != nil {
		log.Fatal(err)
	}

	newfile, err := os.Create("multitemplate-index-preferred.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()

    fmt.Println("------Writing file-b template---------------")
	err = tpl.ExecuteTemplate(newfile,"template-file-b.text",tpl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------Writing file-c template---------------")
	err = tpl.ExecuteTemplate(newfile,"template-file-c.text",tpl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------Writing file-c template---------------")
	err = tpl.ExecuteTemplate(newfile,"template-file-a.text",tpl)
	if err != nil {
		log.Fatal(err)
	}

}