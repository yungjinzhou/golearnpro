package main

import (
	"io"
	"log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
		<input type="text" name="in"/>
		<input type="submit" value="Submit"/>
	</form></html></body>`


type HttpHandler func(http.ResponseWriter, *http.Request)

// SimpleServer handle a simple GET request
func SimpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello world</h1>")
}

// FormServer handle a form , both the GET which displays the form and the POST which process it
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, request.FormValue("in"))
	}
}


func logPanic(function HttpHandler) HttpHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic %v", r.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}


func main() {
	http.HandleFunc("/test1", logPanic(SimpleServer))
	http.HandleFunc("/test2", logPanic(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}