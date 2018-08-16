package signals

import (
	"os"
	"syscall"
	"log"
	"os/signal"
	"net/http"
)

func SetupSignals(sigs chan os.Signal, server *http.Server) {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go CatchSignals(sigs, server)
}

func CatchSignals(signals chan os.Signal, server *http.Server) {
	sig := <-signals
	switch sig {
	case syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL:
		log.Println("Terminate signal received, stopping server...")
		server.Shutdown(nil)
	}
}
