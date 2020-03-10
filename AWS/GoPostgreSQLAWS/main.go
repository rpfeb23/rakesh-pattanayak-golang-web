package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"html/template"
	"strconv"
)

var db *sql.DB
var err error
var inputtologerror string
var tpl *template.Template


type customer struct {
	Customer_id    int
	Customer_fname string
	Customer_lname string
	Customer_age    int
}

func init()  {
	/* DB Instance id : database-1
		   DB Name : rakeshpostgresqldb  (Initial Db Name)
	       DB Master UserId : postgresuser
		   DB Master Password : mypassword
		   END Point : database-1.c0g33tfx1wxa.us-west-1.rds.amazonaws.com
			PORT : 5432
            The ONE created in Default VPC for Testing
	db, err =  sql.Open("postgres", "postgres://postgresuser:mypassword@database-1.c0g33tfx1wxa.us-west-1.rds.amazonaws.com:5432/rakeshpostgresqldb?sslmode=disable")

	*/
	db, err =  sql.Open("postgres", "postgres://postgresuser:mypassword@database-1.c0g33tfx1wxa.us-west-1.rds.amazonaws.com:5432/rakeshpostgresqldb?sslmode=disable")

	inputtologerror = "Open"
	logerror(err, inputtologerror)

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Successful Ping")
	}

	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main()  {
	defer db.Close()

	servemux()

	err = http.ListenAndServe(":80", nil)
	// for test enviornment
	//err = http.ListenAndServe(":8080", nil)
	//
	inputtologerror = "ListenAndServe"
	logerror(err,inputtologerror)
}

func servemux()  {
	http.HandleFunc("/", indexfunc)
	http.HandleFunc("/index", indexfunc)
	http.HandleFunc("/insert", inserttodb)
	http.HandleFunc("/fetchall", fetchall)
	http.HandleFunc("/getcustomer", getcustomer)
	http.HandleFunc("/deletecustomer", deletecustomer)
	http.HandleFunc("/createtable", createtable)
}
func indexfunc(w http.ResponseWriter, r *http.Request)  {

	response, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")

	if err != nil {
		log.Fatal(err)
	}
	instanceid_bs := make([]byte, response.ContentLength)
	response.Body.Read(instanceid_bs)
	response.Body.Close()
	EC2instanceid := string(instanceid_bs)

	//  for Test Environment
	//EC2instanceid := "aaaaa"
	//

	err = tpl.ExecuteTemplate(w,"index.gohtml",EC2instanceid)
	inputtologerror = "ExecuteTemplate index.gohtml"
	logerror(err, inputtologerror)
}

func inserttodb(w http.ResponseWriter, r *http.Request)  {

	if r.Method == http.MethodPost{
		/* Otherway to Read the BODY
		bs := make([]byte, r.ContentLength)
		r.Body.Read(bs)
		body := string(bs)
		fmt.Println("POST Insert Form BODY : ",body)
		*/
		err = r.ParseForm()
		inputtologerror = "ParseForm inserttodb"
		logerror(err, inputtologerror)
		// you can access r.PostForm only after r.ParseForm
		fmt.Println("POST Insert Form Value : ",r.Form, r.PostForm)


		age,_ := strconv.Atoi(r.FormValue("age"))
		c := customer{
			Customer_id:  0  ,
			Customer_fname: r.FormValue("first"),
			Customer_lname: r.FormValue("last"),
			Customer_age :  age ,
		}
		fmt.Println("customer entry : ", c.Customer_fname,c.Customer_lname,c.Customer_age)

		_, err := db.Exec("INSERT INTO customer (customer_fname, customer_lname, customer_age) VALUES ($1,$2,$3)",c.Customer_fname,c.Customer_lname,c.Customer_age)

		inputtologerror = "Exec INSERT "
		logerror(err, inputtologerror)

		/* POSTGRES DOES NOT SUPPORT LastInsertId seems
		customer_id_int64, err := result.LastInsertId()
		c.Customer_id = int(customer_id_int64)
		fmt.Println("Successful INSERT Last ID : ", c.Customer_id)
		*/
		err = tpl.ExecuteTemplate(w,"insert.gohtml",c)
		inputtologerror = "ExecuteTemplate insert.gohtml"
		logerror(err, inputtologerror)
	}else {
		err = tpl.ExecuteTemplate(w,"insert.gohtml",nil)
		inputtologerror = "ExecuteTemplate insert.gohtml"
		logerror(err, inputtologerror)
	}

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

	err = tpl.ExecuteTemplate(w,"fetchall.gohtml",c)
	inputtologerror = "ExecuteTemplate fetchall.gohtml"
	logerror(err, inputtologerror)

}


func getcustomer(w http.ResponseWriter, r *http.Request)  {
	c1 := customer{}
	if r.Method == "GET" {
		/* name= fields from template html */
		cid, _ := strconv.Atoi(r.URL.Query().Get("customerid"))
		age, _ := strconv.Atoi(r.URL.Query().Get("age"))
		c1 = customer{
			Customer_id:    cid,
			Customer_fname: r.URL.Query().Get("first"),
			Customer_lname: r.URL.Query().Get("last"),
			Customer_age:   age,
		}
		log.Println(c1)
	}
	c := []customer{}
	// Serach Only if one of the field is populated
	if (c1.Customer_id > 0) || (c1.Customer_age > 0) || (c1.Customer_lname !="") || (c1.Customer_fname != "") {

		querystring := "select customer_id, customer_fname, customer_lname, customer_age from customer where "
		if (c1.Customer_id > 0) {
			querystring += "customer_id = $1"
		}
		if (c1.Customer_fname != "") {
			querystring += " and customer_fname = $2"
		}
		if (c1.Customer_lname != "") {
			querystring += " and customer_lname = $3"
		}
		if (c1.Customer_age > 0) {
			querystring += " and customer_age = $4"
		}
		fmt.Println(querystring)

		rows, err := db.Query(querystring, c1.Customer_id, c1.Customer_fname, c1.Customer_lname, c1.Customer_age)
		inputtologerror = "Exec select getcustomer "
		logerror(err, inputtologerror)

		defer rows.Close()

		for rows.Next() {
			ci := customer{}
			err = rows.Scan(&ci.Customer_id, &ci.Customer_fname, &ci.Customer_lname, &ci.Customer_age)

			inputtologerror = "Rows Scan getcustomer"
			logerror(err, inputtologerror)
			c = append(c, ci)
		}

		fmt.Println(c)
	}

	err = tpl.ExecuteTemplate(w,"customersearch.gohtml",c1)
	inputtologerror = "ExecuteTemplate customersearch.gohtml"
	logerror(err, inputtologerror)
}


func deletecustomer(w http.ResponseWriter, r *http.Request)  {
	//* Called via GET

	if r.URL.Query().Get("customerid") != "" {
		deletecustomerid,_ := strconv.Atoi(r.URL.Query().Get("customerid"))
		result, err := db.Exec("DELETE FROM customer WHERE customer_id = $1;", deletecustomerid)
		inputtologerror = "Exec INSERT "
		logerror(err, inputtologerror)

		count, _ := result.RowsAffected()
		fmt.Println("Total Records Deleted : ", count)
	}

	err = tpl.ExecuteTemplate(w,"customerdelete.gohtml",r.URL.Query().Get("customerid"))
	inputtologerror = "ExecuteTemplate customerdelete.gohtml"
	logerror(err, inputtologerror)

}

func createtable(w http.ResponseWriter, r *http.Request)  {

		_, err := db.Exec("CREATE TABLE public.customer(customer_id integer NOT NULL GENERATED ALWAYS AS IDENTITY (INCREMENT 1 START 1000 MINVALUE 1000 MAXVALUE 9999 ),customer_fname text,customer_lname text,customer_age integer,PRIMARY KEY (customer_id))WITH (OIDS = FALSE);")
		inputtologerror = "Exec CREATE Table "
		logerror(err, inputtologerror)

}

func logerror(err error, inputtologerror string){
	if err != nil {
		log.Printf(" Error while %v , Error is %v \n",inputtologerror,err)
	}else{
		log.Println("Successful ", inputtologerror)
	}

}