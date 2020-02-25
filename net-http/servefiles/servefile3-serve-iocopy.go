package main

import (
	"io"
	"net/http"
	"os"
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
	file1, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "toby not found ", 404)
		return
	}

	defer file1.Close()

	io.Copy(w,file1)
}