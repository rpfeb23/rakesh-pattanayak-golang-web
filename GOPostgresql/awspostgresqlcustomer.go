package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB
var err error
var inputtologerror string

type customer struct{
	Customer_id    int
	Customer_fname string
	Customer_lname string
	Customer_age    int

}
func init(){
	/* DB Instance id : database-1
	   	   DB Name : rakeshpostgresqldb
	          DB Master UserId : postgresuser
	   	   DB Master Password : mypassword
	   	   EN Point : database-1.cla6njvvwbew.us-east-1.rds.amazonaws.com
	   		PORT : 5432
	*/


	db, err =  sql.Open("postgres", "postgres://postgresuser:mypassword@database-1.c0g33tfx1wxa.us-west-1.rds.amazonaws.com:5432/rakeshpostgresqldb?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("We have opened  Postgres")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Successful Ping")
	}

}

func main()  {
	defer db.Close()

	servemux()

	err = http.ListenAndServe(":8080", nil)

	inputtologerror = "ListenAndServe"
	logerror(err,inputtologerror)

}
func servemux() {
	http.HandleFunc("/fetchall", fetchall)
}
func fetchall(w http.ResponseWriter, r *http.Request)  {
	rows, err := db.Query("select customer_id, customer_fname, customer_lname, customer_age  from customer;")
	inputtologerror = "Fetchall Query "
	logerror(err, inputtologerror)

	defer rows.Close()

	c := []customer{}

	for rows.Next(){
		ci := customer{}
		err = rows.Scan(&ci.Customer_id,&ci.Customer_fname,&ci.Customer_lname,&ci.Customer_age)

		inputtologerror = "Rows Scan "
		logerror(err, inputtologerror)
		c = append(c, ci)
	}

	fmt.Println(c)
}
func logerror(err error, inputtologerror string){
	if err != nil {
		log.Printf(" Error while %v , Error is %v \n",inputtologerror,err)
	}else{
		log.Println("Successful ", inputtologerror)
	}

}