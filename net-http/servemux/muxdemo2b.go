// Implement using Handle and Convert using HandlerFunc
package main

import (
	"fmt"
	"net/http"
)

// FUNCTIONS ARE ALSO A TYPE
// below two functions are of type
//       func(http.ResponseWriter, *http.Request)

func d(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You want Dog?")
}

func c(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w," You want Cat?")
}

func main()  {

	fmt.Printf("Type of Function d : %T\n", d)
	fmt.Printf("Type of Function c : %T\n", c)

	mux := http.NewServeMux()
	// HandlerFunc is a TYPE which has an underlying type
	//           func(http.ResponseWriter, *http.Request)
	// d and c functions are of TYPE
	//           func(http.ResponseWriter, *http.Request)
	// But you can not dirctly assign thos two you have to convert
	//           smilart to " type hotdog int " and " "var i int"

	// Handle Takes Handler Since HandlerFunc implements ServeHTTP, you can pass HandlerFunc

	mux.Handle("/dog/", http.HandlerFunc(d))
	mux.Handle("/cat",  http.HandlerFunc(c))

	http.ListenAndServe(":8080", mux)

}