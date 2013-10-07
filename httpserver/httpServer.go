package main

import (
	"fmt"
	"net/http"
)


func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Got request agent" + req.UserAgent())
	fmt.Fprintf(rw, "{\"sucess\":%s,\"message\":\"%s\"}","true","response from go webserver")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
