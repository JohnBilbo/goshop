package main

import "net/http"

func main() {
	mux = http.NewServeMux
	err := http.ListenAndServe(":8005", nil)

}
