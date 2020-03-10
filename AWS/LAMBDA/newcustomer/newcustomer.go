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



	fmt.Println(request.HTTPMethod, request.QueryStringParameters)
	if  request.HTTPMethod == "POST"{

		c := Response{}

		err = json.Unmarshal([]byte(request.Body),&c)
		if err != nil {
			fmt.Println("Unable to Unmarshal Request Body")
		}
		_, err := db.Exec("INSERT INTO customer (customer_fname, customer_lname, customer_age) VALUES ($1,$2,$3)",c.Customer_fname,c.Customer_lname,c.Customer_age)

		if err != nil {
			fmt.Println("Unable to Execute Query")
		}

		APIResponse := events.APIGatewayProxyResponse{Body: "Successfully Created Customer", StatusCode: 201}
		return APIResponse, nil
	}else {
		APIResponse := events.APIGatewayProxyResponse{Body: request.HTTPMethod, StatusCode: 405}
		return APIResponse, nil
	}


}
