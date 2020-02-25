package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"html/template"
	"log"
	"net/http"

)
var db *sql.DB
var err error
var inputtologerror string
var tpl *template.Template
var EC2instanceid  string
var fetchallpath bool

type customer struct {
	Customer_id    int64
	Customer_fname string
	Customer_lname string
	Customer_age    int
	Insatnce_id    string
}
func init()  {
	tpl = template.Must(template.ParseGlob("Templates/*"))
}

func main()  {

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/", EC2instanceinfo)
	http.HandleFunc("/customerlist",fetchall )
	http.ListenAndServe(":80", nil)

}

func Ping(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Pinged Successfully")
	log.Println("Pinged Successfully")
}

func EC2instanceinfo(w http.ResponseWriter, r *http.Request)  {

	response, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")

	if err != nil {
		log.Fatal(err)
	}
	instanceid_bs := make([]byte, response.ContentLength)
	response.Body.Read(instanceid_bs)
	response.Body.Close()
	EC2instanceid = string(instanceid_bs)
    if (fetchallpath == false) {
		err = tpl.ExecuteTemplate(w, "instanceid.gohtml", EC2instanceid)
		inputtologerror = "ExecuteTemplate instanceid.gohtml"
		logerror(err, inputtologerror)
	}
}

func fetchall(w http.ResponseWriter, r *http.Request)  {

	fetchallpath = true
	EC2instanceinfo(w,r)
	fetchallpath = false

	/* EC2instanceid = "Dummy" */

	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(database-2.cla6njvvwbew.us-east-1.rds.amazonaws.com:3306)/rakeshmysqldb?charset=utf8")
	inputtologerror = "Open"
	logerror(err, inputtologerror)


	err = db.Ping()

	inputtologerror = "PingContext"
	logerror(err,inputtologerror)


	defer db.Close()

	rows, err := db.Query("select customer_id, customer_fname, customer_lname, customer_age  from customer")
	inputtologerror = "Fetchall Query "
	logerror(err, inputtologerror)

	defer rows.Close()

	c := []customer{}
	for rows.Next(){
		ci := customer{}
		ci.Insatnce_id = EC2instanceid
		err = rows.Scan(&ci.Customer_id,&ci.Customer_fname,&ci.Customer_lname,&ci.Customer_age)

		inputtologerror = "Rows Scan "
		logerror(err, inputtologerror)
		c = append(c, ci)
	}

	fmt.Println(c)

	err = tpl.ExecuteTemplate(w,"fetchall.gohtml",c)
	inputtologerror = "ExecuteTemplate fetchall.gohtml"
	logerror(err, inputtologerror)

}

func logerror(err error, inputtologerror string){
	if err != nil {
		log.Printf(" Error while %v , Error is %v \n",inputtologerror,err)
	}else{
		log.Println("Successful ", inputtologerror)
	}

}