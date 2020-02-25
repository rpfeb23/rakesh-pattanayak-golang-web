package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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

	request(conn)
	response(conn)

	defer conn.Close()
}

func request(conn net.Conn)  {
	scaner1 := bufio.NewScanner(conn)
	i := 0
	for scaner1.Scan(){
		line := scaner1.Text()
		fmt.Println(line)
		wordsfromline := strings.Fields(line)
		if i == 0{
			// The very first line from browser localhost:8080 is REQUEST LINE an the first Word is the METHOD. A Sample REQUEST Line looks like : GET /rakesh/patt HTTP/1.1
			fmt.Println("*** Interpreted the HTTP METHOD ***",wordsfromline[0],"******")
			fmt.Println("*** URI of the GET Request :" , wordsfromline[1] )
		}

		if line == ""{
			// Finished Reading Headers. The Broswer client sends 1)REQUEST LINE  2) HEADER Lines 3) BLANK Line 4) REQUEST BODY
			break // done with scanning
		}

		i++
	}
}

func response(conn net.Conn)  {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Rakesh is Awesome</title></head><body><strong>Hello World</strong></body></html>`
	// Prints to Connection (clinet)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // STATUS Line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) // HEADER
	fmt.Fprint(conn, "Content-Type: text/html\r\n") // HEADER
	fmt.Fprint(conn, "\r\n") // BLANK Line
	fmt.Fprint(conn, body)   // RESPONSE BODY
}