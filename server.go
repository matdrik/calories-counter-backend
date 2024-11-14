package calories_counter_backend

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("0.0.0.0:%s", port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(c context.Context) error {
	return s.httpServer.Shutdown(c)
}
