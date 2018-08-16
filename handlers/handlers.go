package handlers

import (
	"fmt"
	"net/http"
)

func SetupHandlers() {
	http.HandleFunc("/", RootHandler)
	// add handlers for each mapped path | app and forward returned JSON to caller
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Root path of server, nothing to see here.")
}
