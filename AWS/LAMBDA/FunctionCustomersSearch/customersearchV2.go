package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
 "github.com/aws/aws-lambda-go/events"
	"log"
)
var db *sql.DB
var rows *sql.Rows
var err error
type Response struct {
	Customer_id       int    `json:"customerid"`
	Customer_fname string `json:"customefirstname"`
	Customer_lname  string `json:"customelastname"`
	Customer_age       int    `json:"customeage"`
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

	fmt.Println(request.HTTPMethod, request.QueryStringParameters,request.PathParameters)
	cid := ""

	if  request.HTTPMethod == "GET"{
		cid,_ = request.PathParameters["customerid"]
	}else {
		APIResponse := events.APIGatewayProxyResponse{Body: request.HTTPMethod, StatusCode: 405}
		return APIResponse, nil
	}
       if (cid != "")  {
		   querystring := "select customer_id, customer_fname, customer_lname, customer_age from customer where customer_id = $1"
		   rows, err = db.Query(querystring, cid)
	   }else {
		   querystring := "select customer_id, customer_fname, customer_lname, customer_age from customer"
		   rows, err = db.Query(querystring)
	   }

		if err != nil {
			fmt.Println("Query failed to execute")
			APIResponse := events.APIGatewayProxyResponse{Body: "Query failed to execute", StatusCode: 500}
			return APIResponse, err
		}
		defer rows.Close()
		rs := []Response{}

		for rows.Next() {
			r := Response{}
			err = rows.Scan(&r.Customer_id, &r.Customer_fname, &r.Customer_lname, &r.Customer_age)
			rs = append(rs,r)
		}
		responsebyteslice, _ := json.Marshal(rs)
		headers := map[string]string{
			"Access-Control-Allow-Origin": "*",
		}
	    APIResponse := events.APIGatewayProxyResponse{Body: string(responsebyteslice), StatusCode: 200, Headers: headers}
	return APIResponse, nil

}