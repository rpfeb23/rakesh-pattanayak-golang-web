package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
func main(){

	//* AWSDBMasterUsername:MasterPassword@tcp(DBEndpoint)/DBName
	//* In my case I had DB Identifier in AWS as database-2
	//* DB Name was rakeshmysqldb

	db, err := sql.Open("mysql", "awsuser:mypassword@tcp(database-2.cla6njvvwbew.us-east-1.rds.amazonaws.com:3306)/rakeshmysqldb?charset=utf8")

	if err != nil {
		fmt.Println(err)
	}else{
		log.Println("Successful DB Open")
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc("/", indexfunc)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func indexfunc(w http.ResponseWriter, r *http.Request)  {
		io.WriteString(w, " Successfully Pinged DB")
}