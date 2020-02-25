package main

import (
	"fmt"
	"net/http"
)

type something int

func (st something) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

	switch r.URL.Path {
	case "/dog":
			fmt.Fprintln(w," You want Dog?")
	case "/cat":
		    fmt.Fprintln(w," You want Cat?")
	}

}

func main()  {

	st1 := something(42)

	http.ListenAndServe(":8080", st1)

}