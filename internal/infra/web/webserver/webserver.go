package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	if _, ok := s.Handlers[path]; !ok {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}

	s.Handlers[path][method] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, methods := range s.Handlers {
		for method, handler := range methods {
			s.Router.Method(method, path, handler)
		}
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%s", s.WebServerPort), s.Router); err != nil {
		panic(err)
	}
}
