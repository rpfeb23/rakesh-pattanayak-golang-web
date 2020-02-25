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
		fmt.Println("Connection Established...")
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
		bs := []byte(ln)
		bs = rot13encrypt(bs)
		fmt.Fprintf(conn," Input : %s  Output : %s \n", ln, string(bs)) // Prints on Client
		fmt.Println(string(bs)) // Prints on Server
	}
	defer conn.Close()
}

func rot13encrypt(bs []byte)  []byte{
	var encryptedbs = make([]byte,len(bs))
	for i, v := range bs{
		encryptedbs[i] = v + 13
	}
	return encryptedbs
}