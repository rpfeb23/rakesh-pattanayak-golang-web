
package beginners

import (
	"io"
	"log"
	"os"
	"strings"
)

func main()  {
	name := "James Bond"
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello Bond!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	f , err := os.Create("index1.html")
	if err != nil {
		log.Fatal(" Error Creating index1.html")
	}
	defer f.Close()

	io.Copy(f,strings.NewReader(tpl))
}