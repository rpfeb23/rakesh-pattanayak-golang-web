/*   net.Listen
     net.Accept
         conn implements Reader and Writer Interface
	 net.Close [Close the Listenre and Connection Both]

 WRITE TO CLIENT

 */

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main()  {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	defer li.Close()

	for {

		conn, err := li.Accept()
		fmt.Println("Someone hit the telenet",conn) // Prints on Server
		if err != nil {
			log.Fatal(err)
		}
		// Writes on the Client Terminal

		io.WriteString(conn," Hello From Rakesh's Server")
		fmt.Fprintln(conn,time.Now().Format(time.RFC3339),"How are you today?")
		fmt.Fprintf(conn,"%v","And, How can I help?")

		conn.Close()
	}


}