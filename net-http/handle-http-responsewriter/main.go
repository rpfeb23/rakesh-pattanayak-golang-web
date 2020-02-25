package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// ResponseWriter is an Interface | Request is a Struct
func (hd hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Rakesh ", "Nice Name")
	fmt.Fprintln(w,"From The Server I am saying Hi")
}

func main()  {
	hd1 := hotdog(42)

	http.ListenAndServe(":8080", hd1)

}