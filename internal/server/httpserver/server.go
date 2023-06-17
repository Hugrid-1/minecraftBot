package httpserver

import (
	"context"
	"fmt"
	"github.com/Hugrid-1/minecraftBot/config"
	"net/http"
	"time"
)

type HTTPServer struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func NewHTTPServer(serverConfig config.ServerSettings, handler http.Handler) *HTTPServer {
	address := fmt.Sprintf(":%s", serverConfig.Port)
	server := &http.Server{
		Addr:    address,
		Handler: handler,
	}
	httpServer := &HTTPServer{
		server:          server,
		notify:          make(chan error, 1),
		shutdownTimeout: 100,
	}
	httpServer.start()

	return httpServer
}

func (s *HTTPServer) start() {
	fmt.Printf("[LOG %v]HTTP Server started at %v\n", time.Now().Format("15:04:05"), s.server.Addr)
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *HTTPServer) Notify() <-chan error {
	return s.notify
}

func (s *HTTPServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
