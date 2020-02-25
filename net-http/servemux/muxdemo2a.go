// Implement using HandleFunc
package main

import (
	"fmt"
	"net/http"
)

type dog int

func (d dog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You want Dog?")

}

type cat int

func (c cat) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You want Cat?")

}

func main()  {

	d1 := dog(42)
	c1 := cat(84)
	// ServeMux implements ServeHTTP(Response *Request) so it any value of type ServeMux is also Type Handler. Now we can pass ServeMux to ListenAndServe

	mux := http.NewServeMux()
	mux.HandleFunc("/dog/", d1.ServeHTTP)
	mux.HandleFunc("/cat", c1.ServeHTTP)

	http.ListenAndServe(":8080", mux)

}