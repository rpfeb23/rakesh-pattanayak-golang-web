package main

import (
	"fmt"
	"net/http"
)

func main()  {
	http.HandleFunc("/",indexfunc)
	http.ListenAndServe(":80", nil)
}

func indexfunc(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Hello : We are running inside DOCKER Container")
}