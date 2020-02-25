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
	// ServeMux implements ServeHTTP(Response *Request) so it any value of type ServeMux is also Type Handler. Now we can pass mux which is a Pointer to ServeMux as HANDLER to ListenAndServe

	mux := http.NewServeMux()
	mux.Handle("/dog/", d1) //dog/* anything after dog will be catched
	mux.Handle("/cat", c1)  // any request /cat/** will return 404
	// you have to pass /cat to invoke this handler

	http.ListenAndServe(":8080", mux)

}