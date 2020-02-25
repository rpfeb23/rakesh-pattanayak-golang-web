package main

import (
	"io"
	"net/http"
)

func main()  {

	http.Handle("/", http.StripPrefix("/dogs", http.FileServer(http.Dir("./resources"))))

	http.HandleFunc("/toby",toby)

	http.ListenAndServe(":8080",nil)
}

func toby(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	// When this gets written it call http.Handle("/".....)
	io.WriteString(w,`<img src="/dogs/toby.jpg">`)
}