package main

import (
	"fmt"
	"log"
	"text/template"
	"net/http"
)

type somevariable int

func (sv somevariable) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	//ParseForm populates r.Form and r.PostForm.
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("GET and POST INFO ", r.Form)     // map[string][string]
	fmt.Println(" POST Method Info :",r.PostForm) // map[string[string]
	fmt.Println(" REQUEST HOST : ", r.Host)
	fmt.Println(" FORM Value of Key fname : ", r.FormValue("fname"))
	/*******   For Own Understanding of r.Form **********/
	/*
	fmt.Println("**************** Request INFO************************")
	fmt.Println(*r)
	fmt.Println("****** REQUEST METHOD : ", r.Method)
	fmt.Println("****** REQUEST URI : ", r.URL)

	fmt.Println("GET and POST INFO ", r.Form)     // map[string][string]
	fmt.Println(" POST Method Info :",r.PostForm) // map[string[string]

	for i, v := range r.Form{
		fmt.Print("Index : ",i)
		fmt.Println("Value : ", v)
		for _, v1 := range v{
			fmt.Printf("\t\t\t Individual Values : %v\n", v1)
		}
	}
	*/
	/*******   For Own Understanding of r.Form **********/

	tpl.ExecuteTemplate(w,"index.gohtml",r.Form)

}

var tpl *template.Template
func init()  {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main()  {
	sv1 := somevariable(42)
	err := http.ListenAndServe(":8080",sv1)
	if err != nil {
		fmt.Println(err)
	}
}