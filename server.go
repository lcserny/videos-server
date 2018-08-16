package main

import (
	"net/http"
	"log"
	"flag"
	"github.com/lcserny/videos-server/handlers"
	"github.com/lcserny/videos-server/signals"
	"os"
)

func processAddress() *string {
	address := flag.String("addr", ":8080", "Address and port for the server")
	flag.Parse()
	return address
}

func main() {
	address := processAddress()
	sigs := make(chan os.Signal)
	server := &http.Server{Addr:*address}

	handlers.SetupHandlers()
	signals.SetupSignals(sigs, server)

	log.Printf("Starting webserver on address %s\n", *address)
	log.Println("Press ctrl+c to stop...")
	log.Fatal(server.ListenAndServe())
}
