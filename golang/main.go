package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

var port = "8080"

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body) // игнорируем ошибку!
	fmt.Println("Got request:", string(body))
	fmt.Fprintf(w, "OK")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started")
	http.ListenAndServe(":"+port, nil) // не обрабатываем ошибку
}
