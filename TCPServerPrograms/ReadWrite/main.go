
package main

import (
	"bufio"
	"fmt"
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
		if err != nil {
			log.Fatal(err)
		}
		go handlescan(conn)
	}


}

// Type of conn is net.Conn
func handlescan(conn net.Conn)  {

	err := conn.SetReadDeadline(time.Now().Add(15 * time.Second)) // if you keep only 10 it will be nanosecond
	if err != nil {
		fmt.Println("Connection TimeOut")
	}
	fmt.Println(conn)

	// conn is a Reader as well as Writer (Like File)
	scaner1 := bufio.NewScanner(conn)
	for scaner1.Scan(){
		ln := scaner1.Text()
		fmt.Println(ln)  // Prints on Server
		fmt.Fprintln(conn,"I Heard You Say : ", ln) // Writes on to the Connection which will be displayed on Client
	}
	defer conn.Close()

	fmt.Println("Code GOT here......")
}