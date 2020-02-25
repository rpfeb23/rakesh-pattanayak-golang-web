package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
)

var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseGlob("template1.html"))
}

type person struct {
	Name string
	Age int
}

func (p person) Personfunc1()  string {
	return fmt.Sprint("From personfunc1")
}

func (p person) Personfunc2(name string) string {
	return fmt.Sprintf("Hello , %v", name)
}

func (p person) Personfunc3(age int, stringfromfunc2 string)  string {
	returnstring := "Your Age is " + strconv.Itoa(age) + " " + stringfromfunc2
	return returnstring
}

func main()  {

	inputstruct := person{
		"James Bond",
		40,
	}

	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",inputstruct)

	if err != nil {
		log.Fatal(err)
	}


}