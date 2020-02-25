package main

import (
	"fmt"
	"net/http"
)

func dogfunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You are a Dog")
}

func myfunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," Hi Rakesh")
}

func indexfunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," Hello. Welcome to Earth")
}

func main()  {

	http.HandleFunc("/", indexfunc)

	http.HandleFunc("/me",myfunc)

	http.HandleFunc("/dog",dogfunc)

	http.ListenAndServe(":8080",nil)

}