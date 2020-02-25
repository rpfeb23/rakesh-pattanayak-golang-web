// go run concatemain.go > inex.html
package beginners

import "fmt"

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
	fmt.Println(tpl)
}