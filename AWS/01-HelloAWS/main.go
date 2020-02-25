package	main

import (
	"fmt"
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/", indexFunc)
	http.ListenAndServe(":80",nil)
}

func indexFunc(w http.ResponseWriter, r *http.Request)  {
	response, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")

	if err != nil {
		log.Fatal(err)
	}
	instanceid_bs := make([]byte, response.ContentLength)
	response.Body.Read(instanceid_bs)
	response.Body.Close()
	EC2instanceid := string(instanceid_bs)

	fmt.Fprintln(w,"Hello Rakesh. You are now on AWS Instance ", EC2instanceid)
}