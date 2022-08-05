package main

import (
	"fmt"
	"log"
	"net/http"
)

func TestHttpServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer Handler")
	//fmt.Fprint(w, "Hello, "+req.URL.Path[1:])

}


func main() {
	http.HandleFunc("/", TestHttpServer)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer", err.Error())

	}
}

