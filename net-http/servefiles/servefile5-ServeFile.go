package main

import (
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/", indexfunc)
	http.HandleFunc("/toby.jpg", toby)
	http.ListenAndServe(":8080",nil)

}

func indexfunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	io.WriteString(w, `
	<!--Serving from our server-->
	<img src="/toby.jpg">
	`)
}

func toby(w http.ResponseWriter, r *http.Request)  {
	http.ServeFile(w,r,"toby.jpg")
}