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

	http.Handle("/", http.HandlerFunc(indexfunc))

	http.Handle("/me",http.HandlerFunc(myfunc))

	http.Handle("/dog",http.HandlerFunc(dogfunc))

	http.ListenAndServe(":8080",nil)

}