package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/harunsasmaz/kubernetes-playground/internal/service/hello"
)

func main() {
	server := hello.NewServer()

	errCh := make(chan error)

	go func() {
		if err := server.Serve(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()
	log.Printf("Starting server at %s\n", server.Addr)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	select {
	case err, ok := <-errCh:
		if ok {
			log.Println("server error:", err)
		}

	case sig := <-sigCh:
		log.Printf("Signal %s received\n", sig)
		if err := server.Shutdown(); err != nil {
			log.Println("Failed to shutdown server:", err)
		}
		log.Println("server shutdown")
	}
}
