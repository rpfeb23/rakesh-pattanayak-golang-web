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

	handlerequest(conn)

	defer conn.Close()
}

func handlerequest(conn net.Conn)  {
	scaner1 := bufio.NewScanner(conn)
	i := 0
	for scaner1.Scan(){
		line := scaner1.Text()
		fmt.Println(line)
		wordsfromline := strings.Fields(line)
		if i == 0{
			httpmethod := strings.Fields(line)[0]
			httpURI := wordsfromline[1]
			fmt.Println("*** HTTP METHOD ***",httpmethod)
			fmt.Println("*** URI of Request :" , httpURI )
			muxrouterequest(conn, httpmethod, httpURI)
		}

		if line == ""{
			// Finished Reading Headers. The Broswer client sends 1)REQUEST LINE  2) HEADER Lines 3) BLANK Line 4) REQUEST BODY
			break // done with scanning
		}
		i++
	}
}

func muxrouterequest(conn net.Conn, httpmethod, httpURI string)  {

	if httpmethod == "GET" && httpURI == "/" {
		indexresponse(conn)
	}

	if httpmethod == "GET" && httpURI == "/contact"{
		contactresponse(conn)
	}

	if httpmethod == "GET" && httpURI == "/apply" {
		getapplyresponse(conn)
	}

	if httpmethod == "POST" && httpURI == "/apply" {
		postapplyresponse(conn)
	}
}

func indexresponse(conn net.Conn)  {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>INDEX</strong><br>
	<a href="/">index</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	// Prints to Connection (clinet)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // STATUS Line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body)) // HEADER
	fmt.Fprint(conn, "Content-Type: text/html\r\n") // HEADER
	fmt.Fprint(conn, "\r\n") // BLANK Line
	fmt.Fprint(conn, body)   // RESPONSE BODY
}

func contactresponse(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>CONTACT</strong><br>
	<a href="/">index</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func getapplyresponse(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>
	<strong>APPLY</strong><br>
	<a href="/">index</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	<form method="POST" action="/apply">
	<input type="submit" value="apply">
	</form>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func postapplyresponse(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>APPLY PROCESS</strong><br>
	<a href="/">index</a><br>
	<a href="/contact">contact</a><br>
	<a href="/apply">apply</a><br>
	</body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}