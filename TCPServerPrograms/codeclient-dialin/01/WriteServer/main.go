package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main()  {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\n Hello from TCP Server \n")
		fmt.Println(conn, "I hope it is great") // Writes to Server
		fmt.Fprintln(conn,"I hope it is great") // Writes to Client

		conn.Close()
	}

}