package main

import (
	"text/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}
var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseGlob("*.html"))
}
func main()  {
    xs := []string{"A", "B", "C", "D"}

    //Demonstrate Index
	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",xs)
	if err != nil {
		log.Fatal(err)
	}

	person1 := person{Name: "Rakesh", Age:35}
	person2 := person{Name: "Rajesh", Age:38}
	person3 := person{
		Name: "",
		Age:  40,
	}
	person4 := person{}

	people := []person{person1 , person2, person3,person4}

	err = tpl.ExecuteTemplate(os.Stdout,"template2.html",people)
	if err != nil {
		log.Fatal(err)
	}


	err = tpl.ExecuteTemplate(os.Stdout,"template3.html",people)
	if err != nil {
		log.Fatal(err)
	}

}