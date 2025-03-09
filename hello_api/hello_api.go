package main

import "net/http"

// type dataStore struct {
// 	users map[string]string
// }

func hello(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
