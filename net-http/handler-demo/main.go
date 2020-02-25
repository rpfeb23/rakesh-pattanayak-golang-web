package main

import (
	"fmt"
	"net/http"
)

type somevariable int
var sv somevariable

// It Implements Handler Interface so any traffic
//   http://locallost:8080
//   URI after 8080 will come here
//         eg. localhost:8080/
//             localhost:8080/abcd
//             localhost:8080/abcd/xyz
//   All of the above will come to this function to be served
func (sv somevariable) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" Any Value of Type somevariable is of Type INT and Type Handler ") // Gets Printed on Server

	fmt.Fprintln(w," Hello from Server ", sv) // Write to Client
}

func main()  {
	sv1 := somevariable(42)
	fmt.Println("Value of sv1 : ", sv1)
	fmt.Printf("Type of sv1 : %T\n", sv1)

	sv = 42
	fmt.Println("Value of sv : ", sv)
	fmt.Printf("Type of sv : %T\n", sv)


	http.ListenAndServe(":8080", sv1)


}