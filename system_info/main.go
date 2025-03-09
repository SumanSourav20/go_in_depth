package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("hello %s", r.Host)))
	})
	http.ListenAndServe("localhost:8080", nil)
}
