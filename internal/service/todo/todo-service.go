package todo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/harunsasmaz/kubernetes-playground/internal/handler"
)

type Server struct {
	handler   *handler.Handler
	router    *mux.Router
	Addr      string
	MustClose bool
	server    *http.Server
}

func (s *Server) InitializeRoutes() {

	router := mux.NewRouter().PathPrefix("/todos").Subrouter()

	router.HandleFunc("/health", s.handler.Health).Methods("GET")
	router.HandleFunc("", s.handler.GetAll).Methods("GET")
	router.HandleFunc("/done", s.handler.GetDone).Methods("GET")
	router.HandleFunc("/waiting", s.handler.GetRemaining).Methods("GET")
	router.HandleFunc("/{id}", s.handler.Get).Methods("GET")
	router.HandleFunc("", s.handler.Create).Methods("POST")
	router.HandleFunc("/status", s.handler.UpdateStatus).Methods("PUT")
	router.HandleFunc("/{id}", s.handler.Delete).Methods("DELETE")

	s.router = router
}

func NewServer() *Server {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		log.Fatalf("failed to retrieve port from env")
	}

	server := &Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
		handler: handler.New(context.Background()),
	}

	return server
}

func (s *Server) Serve() error {
	s.InitializeRoutes()

	s.server = &http.Server{
		Addr:    s.Addr,
		Handler: s.router,
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
