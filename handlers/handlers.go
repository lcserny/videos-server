package handlers

import (
	"net/http"
	"fmt"
)

func SetupHandlers() {
	http.HandleFunc("/", RootHandler)
}

func RootHandler(writer http.ResponseWriter, request *http.Request)  {
	fmt.Fprintln(writer, "Root path of server, nothing to see here.")
}