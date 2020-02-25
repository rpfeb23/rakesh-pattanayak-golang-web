package main

import "net/http"

func main()  {
	// Another way to serve all of your files at Root
	// Exception if you have index.html it will take precedence and will be displayed

	// try changing the index1.html to index.html and see difference
	
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}