package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type famouspersons struct{
	Fname string
	Lname string
	Age int
}
// Function Map functions can return max 2 returns out of which 1 has to be error.
var fm =template.FuncMap {
	"name1" : func1,
	"name2" : func2,
}
// returns birth year
func func1(age int) int {
    currentyear := time.Now().Year()
    birthyear := currentyear - age
    return birthyear
}
// Return first 3 bytes of string
func func2(lname,fname string) (string, error) {
	lname = lname[0:3]
	returnname := fname + " " + lname
	return returnname, nil
}
var tpl *template.Template
func init()  {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("template1.html"))
}

func main()  {

	fp1 := famouspersons{
		Fname: "Mahatma",
		Lname: "Gandhi",
		Age:   100,
	}

	fp2 := famouspersons{
		Fname: "Martin",
		Lname: "Luther King",
		Age:   120,
	}

	fps := []famouspersons{fp1,fp2}

	err := tpl.ExecuteTemplate(os.Stdout,"template1.html",fps)
	if err != nil {
		log.Fatal(err)
	}

}