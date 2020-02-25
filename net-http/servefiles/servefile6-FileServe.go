package main

import (
	"io"
	"net/http"
)

func main()  {
	// Serve the Entire Directory
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog",toby)
	http.ListenAndServe(":8080",nil)
}

func toby(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	// When this gets written it call http.Handle("/".....)
	// and it will serve /toby.jpg file form the directory
	io.WriteString(w,`<img src="toby.jpg">`)
}