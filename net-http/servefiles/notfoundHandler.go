package main

import (
	"fmt"
	"net/http"
)

func main()  {

	http.HandleFunc("/", mainpage)

	// Chrome sends /favico.ico upon each refresh Other browsers dont
    // Try commenting this and see Terminal
	//http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
	
}

func mainpage(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL)
	fmt.Fprintln(w, " Check your Terminal ")
}
