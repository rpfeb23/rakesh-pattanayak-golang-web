//Using DefaultServeMux from http Package No need to create our own ServeMux and handle the URI using http.HandleFunc

// Look the signature of HandleFunc it takes a Function with argument similar to ServeHTTP

package main

import (
	"fmt"
	"net/http"
)

func func1(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You want Dog?")
}
func func2(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You want Cat?")
}

func main()  {

	http.HandleFunc("/dog/",func1)
	http.HandleFunc("/cat",func2 )

	http.ListenAndServe(":8080", nil) // Passed nil for DefaultServeMux from http package
}