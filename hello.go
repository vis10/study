// test
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, welcome to this page. %s!", r.URL.Path[1:])
	fmt.Println("hit")
}

func main() {
    http.HandleFunc("/", handler)
    err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}