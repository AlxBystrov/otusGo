package app

import (
	"context"
	"time"

	storage "github.com/AlxBystrov/otusGo/hw12_13_14_15_calendar/internal/storage"
	// "github.com/google/uuid"
)

type App struct {
	storage Storage
	logger  Logger
	host    string
	port    int
}

type Logger interface {
	Debug(string, ...any)
	Info(string, ...any)
	Warning(string, ...any)
	Error(string, ...any)
}

type Storage interface {
	CreateEvent(event storage.Event) (storage.Event, error)
	UpdateEvent(id string, event storage.Event) (storage.Event, error)
	DeleteEvent(id string) error
	GetEventsDay(day time.Time) []storage.Event
	GetEventsWeek(day time.Time) []storage.Event
	GetEventsMonth(day time.Time) []storage.Event
}

func New(logger Logger, storage Storage) *App {
	return &App{
		logger:  logger,
		storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {

	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}
