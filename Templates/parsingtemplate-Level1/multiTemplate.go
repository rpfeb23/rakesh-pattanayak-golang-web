// Single Template File using
//     1. Add first template using
//        template.ParseFiles (template is Package) retruns pointer
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

func main()  {
	tpl,err := template.ParseFiles("template-file-a.text","template-file-c.text")
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("template-file-b.text")

	if err != nil {
		log.Fatal(err)
	}

	// Ranomly executes any template even though tpl has reference to all 3 template files
	err = tpl.Execute(os.Stdout,tpl)
	if err != nil {
		log.Fatal(err)
	}

	newfile, err := os.Create("multitemplate-index-option1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer newfile.Close()

    fmt.Println("------------------------------------------------")
	err = tpl.ExecuteTemplate(newfile,"template-file-b.text",tpl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------------------------")
	err = tpl.ExecuteTemplate(newfile,"template-file-c.text",tpl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("------------------------------------------------")
	err = tpl.ExecuteTemplate(newfile,"template-file-a.text",tpl)
	if err != nil {
		log.Fatal(err)
	}

}