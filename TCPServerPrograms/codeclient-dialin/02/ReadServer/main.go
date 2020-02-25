
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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

func handlescan(conn net.Conn)  {

	scaner1 := bufio.NewScanner(conn)
	for scaner1.Scan(){
		ln := scaner1.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	fmt.Println("Code GOT here......")
}