package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	http.HandleFunc("/", processmain)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080",nil)
}

func processmain(w http.ResponseWriter, r *http.Request)  {

	var filebytestream []byte
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("uploadfile1")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(" Multi Part File :", f)
		fmt.Println(" File Header     :", h)
		defer f.Close()

		filebytestream, err = ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(filebytestream)

		nf, err := os.Create("NewFile.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		_, err = nf.Write(filebytestream)

		if err != nil {
			fmt.Println(err)
		}

	}


	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="uploadfile1">
	<input type="submit">
	</form>
	<br>`+string(filebytestream))

}