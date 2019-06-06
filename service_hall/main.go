package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rec hello router2")
	io.WriteString(w, "Hello world!")
}

func main() {
	go func() {
		http.HandleFunc("/", hello)
		http.ListenAndServe(":8000", nil)
	}()
	time.Sleep(time.Minute)
	fmt.Println("service time end")
}
