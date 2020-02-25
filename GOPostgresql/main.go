package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"log"
)

var db *sql.DB

type employee struct{
	Employee_id int
	Employee_name string
	Employee_age int
	Employee_salary string

}

func main()  {

	/* Assumes you have created a DB `company'.
	   you have created user 'rakesh'

	   CREATE USER rakesh with password 'password';
	   grant all privileges on database company to rakesh;
	   grant all privileges on table employees to rakesh;

	   or you can make 'rakesh' superuser with commanf
	   ALTER user rakesh with superuser;
	 */
	db, err := sql.Open("postgres","postgres://rakesh:password@localhost/company?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("We have opened  Postgres")
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Successful Ping")
	}



	rows, err := db.Query("SELECT * FROM employees;")
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Successful Query")
	}
	defer rows.Close()
	e := make([]employee,0)

	for rows.Next(){
		e1 := employee{}
		err = rows.Scan(&e1.Employee_id,&e1.Employee_name,&e1.Employee_age,&e1.Employee_salary)
		if err != nil {
			fmt.Println(err)
		}
		e = append(e,e1)
	}

	//Print what you got

	for i,v := range e{
		fmt.Printf("At Index %v the employee table has value %v \n", i,v)
	}

}