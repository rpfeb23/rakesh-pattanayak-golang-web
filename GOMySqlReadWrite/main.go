package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	Customer_id    int64
	Customer_fname string
	Customer_lname string
	Customer_age    int
}

func init()  {

	tpl = template.Must(template.ParseGlob("Templates/*"))
}

func main()  {
	db, err = sql.Open("mysql", "awsuser:mypassword@tcp(database-2.cla6njvvwbew.us-east-1.rds.amazonaws.com:3306)/rakeshmysqldb?charset=utf8")
	inputtologerror = "Open"
	logerror(err, inputtologerror)

	ping()

	defer db.Close()

	servemux()
	err = http.ListenAndServe(":8080", nil)
	inputtologerror = "ListenAndServe"
	logerror(err,inputtologerror)
}

func ping()  {
	err = db.Ping()

	inputtologerror = "PingContext"
	logerror(err,inputtologerror)
}

func servemux()  {
	http.HandleFunc("/", indexfunc)
	http.HandleFunc("/index", indexfunc)
	http.HandleFunc("/insert", inserttodb)
	http.HandleFunc("/fetchall", fetchall)
	http.HandleFunc("/getcustomer", getcustomer)
	http.HandleFunc("/deletecustomer", deletecustomer)
}
func indexfunc(w http.ResponseWriter, r *http.Request)  {

	err = tpl.ExecuteTemplate(w,"index.gohtml",nil)
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
		fmt.Println("customer entry : ", c)


		stmt, err := db.Prepare("INSERT INTO customer (customer_fname, customer_lname, customer_age) VALUES (?,?,?)")

		inputtologerror = "Prepare INSERT "
		logerror(err, inputtologerror)

		result, err := stmt.Exec(c.Customer_fname,c.Customer_lname,c.Customer_age)
		inputtologerror = "Exec INSERT "
		logerror(err, inputtologerror)

		c.Customer_id, err = result.LastInsertId()
		fmt.Println("Successful INSERT Last ID : ", c.Customer_id)

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
	rows, err := db.Query("select customer_id, customer_fname, customer_lname, customer_age  from customer")
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
			Customer_id:    int64(cid),
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
			querystring += "customer_id = ?"
		}
		if (c1.Customer_fname != "") {
			querystring += " and customer_fname = ?"
		}
		if (c1.Customer_lname != "") {
			querystring += " and customer_lname = ?"
		}
		if (c1.Customer_age > 0) {
			querystring += " and customer_age = ?"
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

		stmt, err := db.Prepare("DELETE FROM customer WHERE customer_id = ?")

		inputtologerror = "Prepare DELETE "
		logerror(err, inputtologerror)

		result, err := stmt.Exec(r.URL.Query().Get("customerid"))
		inputtologerror = "Exec INSERT "
		logerror(err, inputtologerror)

		count, _ := result.RowsAffected()
		fmt.Println("Total Records Deleted : ", count)
	}

	err = tpl.ExecuteTemplate(w,"customerdelete.gohtml",r.URL.Query().Get("customerid"))
	inputtologerror = "ExecuteTemplate customerdelete.gohtml"
	logerror(err, inputtologerror)

}

func logerror(err error, inputtologerror string){
	if err != nil {
		log.Printf(" Error while %v , Error is %v \n",inputtologerror,err)
	}else{
		log.Println("Successful ", inputtologerror)
	}

}