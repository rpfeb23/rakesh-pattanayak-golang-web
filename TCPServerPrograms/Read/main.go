/*   net.Listen
     net.Accept
         conn implements Reader and Writer Interface
	 net.Close

OPEN a Webbrowser and hit loclahost:8080 to Read from CLIENT
 */

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

// Type of conn is net.Conn
func handlescan(conn net.Conn)  {

	fmt.Println(conn)

	// conn is a Reader as well as Writer (Like File)
	scaner1 := bufio.NewScanner(conn)
	for scaner1.Scan(){
		ln := scaner1.Text()
		fmt.Println(ln)  //Prints on Server
	}
	defer conn.Close()

	fmt.Println("Code GOT here......")
}