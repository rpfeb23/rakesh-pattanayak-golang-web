// this will not be served as we have not set it up
package main

import (
	"io"
	"net/http"
)

func main()  {
	http.HandleFunc("/", indexfunc)
	http.ListenAndServe(":8080",nil)

}

func indexfunc(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","text/html; charset=utf-8")

	io.WriteString(w, `
	<!--Image does not serve-->
	<img src="/toby.jpg">
	`)
}