package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
 "github.com/aws/aws-lambda-go/events"
	"log"
	"strconv"
)
var db *sql.DB
var err error
type Request struct {
	Customer_id       int    `json:"customerid"`
	Customer_fname string `json:"customefirstname"`
	Customer_lname  string `json:"customelastname"`
	Customer_age       int    `json:"customeage"`
}

type Response struct {
	Title string `json: "title"`
	Message string `json: "message"`
	Status  int `json: "status"`
	Responseresource string `json:"responseresource"`
}

func main()  {
	lambda.Start(Handler)
}
func Handler(request events.APIGatewayProxyRequest)  (events.APIGatewayProxyResponse, error){

	db, err =  sql.Open("postgres", "postgres://postgresuser:mypassword@database-1.cla6njvvwbew.us-east-1.rds.amazonaws.com:5432/rakeshpostgresqldb?sslmode=disable")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}else {
		fmt.Println("Successful Ping")
	}
	fmt.Println("hello Rakesh")
	defer db.Close()


	fmt.Println(request.HTTPMethod, request.QueryStringParameters)
	if  request.HTTPMethod == "POST"{

		c := Request{}
		r := Response{}

		err = json.Unmarshal([]byte(request.Body),&c)
		if err != nil {
			fmt.Println("Unable to Unmarshal Request Body",request.Body)
		}
		_, err := db.Exec("INSERT INTO customer (customer_fname, customer_lname, customer_age) VALUES ($1,$2,$3)",c.Customer_fname,c.Customer_lname,c.Customer_age)

		if err != nil {
			fmt.Println("Unable to Execute Query")
			r.Title = "Failure"
			r.Message = "Unable to Execute Query"
			r.Status = 503
		}else{
			r.Title = "Success"
			r.Message = "Successfully Created Customer"
			r.Status = 201
		}

		rows, err := db.Query("SELECT currval(pg_get_serial_sequence('customer','customer_id'))")
		if err != nil {
			fmt.Println("Query failed to execute")
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&c.Customer_id)
		}

		r.Responseresource = "https://6r9gxm28i2.execute-api.us-east-1.amazonaws.com/Dev/customers/"+strconv.Itoa(c.Customer_id)

		headers := map[string]string{
			"Access-Control-Allow-Origin": "*",
		}

		responsebyteslice, err := json.Marshal(r)
		if err != nil{
			fmt.Println(err)
		}

		APIResponse := events.APIGatewayProxyResponse{Body: string(responsebyteslice) , StatusCode: r.Status, Headers: headers}
		return APIResponse, nil

	}else {
		APIResponse := events.APIGatewayProxyResponse{Body: request.HTTPMethod, StatusCode: 405}
		return APIResponse, nil
	}
}
