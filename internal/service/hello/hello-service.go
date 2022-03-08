package hello

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	Addr      string
	MustClose bool
	server    *http.Server
}

func NewServer() *Server {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatalf("failed to retrieve port from env")
	}

	server := &Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", port),
	}

	return server
}

func (s *Server) Serve() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting", s.GreetingHandler)
	mux.HandleFunc("/health", s.HealthHandler)

	s.server = &http.Server{
		Addr:    s.Addr,
		Handler: mux,
	}

	if err := s.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatalf(err.Error())
		}
	}

	return nil
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return s.ShutdownWithContext(ctx)
}

func (s *Server) ShutdownWithContext(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		if s.MustClose {
			s.server.Close()
		}
		return err
	}

	return nil
}

func (s *Server) GreetingHandler(r http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		r.WriteHeader(http.StatusNotFound)
		return
	}

	r.WriteHeader(http.StatusOK)
	r.Write([]byte("Hello World"))
}

func (s *Server) HealthHandler(r http.ResponseWriter, req *http.Request) {
	r.WriteHeader(http.StatusOK)
	r.Write([]byte("Healthy!"))
}
