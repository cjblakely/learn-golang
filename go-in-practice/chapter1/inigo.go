package main

import (
	"net/http"
	"fmt"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Inigo Montoya")
}

func main() {
	//visit http://localhost:4000/
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:4000", nil)
}