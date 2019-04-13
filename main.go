package main

import (
	"net/http"
)

func main() {

	//http.ListenAndServe(":8080", http.HandlerFunc(WsHandler))
	http.ListenAndServe(":8090", http.HandlerFunc(WsGenerator))

}
