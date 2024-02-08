package internalhttp

import (
	"context"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/AlxBystrov/otusGo/hw12_13_14_15_calendar/internal/storage"
)

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warning(string, ...any)
	Error(string, ...any)
}

type Application interface { // TODO
	CreateEvent(ctx context.Context, title, description string, date time.Time, duration, notifyBefore time.Duration, userID int) (storage.Event, error)
}
type Server struct {
	logger Logger
	app    Application
	srv    http.Server
}

func NewServer(logger Logger, app Application, host string, port int) *Server {
	return &Server{
		logger: logger,
		app:    app,
		srv: http.Server{
			Addr: net.JoinHostPort(host, strconv.Itoa(port)),
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	http.HandleFunc("/health", loggingMiddleware(s.Health))
	http.HandleFunc("/createEvent", loggingMiddleware(s.CreateEvent))

	if err := s.srv.ListenAndServe(); err != nil {
		s.logger.Error("HTTP server failed", "error", err)
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}

func (s *Server) Health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "health is ok")
}

func (s *Server) CreateEvent(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "{\"data\": \"ok\"}")
}
