package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hello World!")
	return
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}