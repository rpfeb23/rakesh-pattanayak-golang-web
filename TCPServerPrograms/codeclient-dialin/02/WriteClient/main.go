package main

import (
	"fmt"
	"net"
)

func main()  {
	conn, err := net.Dial("tcp","localhost:8080")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, " I am from Client. Do you recognise me?")

}