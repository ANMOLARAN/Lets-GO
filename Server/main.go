package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("yes you are correct"))
}
func customHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to custome endpoint")
}
func main() {

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/custom", customHandler)

	fmt.Println("Starting port at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting at port 8080",err)
	}

}
