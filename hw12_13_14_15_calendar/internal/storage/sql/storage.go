package sqlstorage

import (
	"context"
	"fmt"
	// "database/sql"
	"time"

	"github.com/AlxBystrov/otusGo/hw12_13_14_15_calendar/internal/storage"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	host     string
	port     int
	database string
	user     string
	password string
	conn     *pgx.Conn
}

func New(host string, port int, database string, user string, password string) *Storage {
	return &Storage{
		host:     host,
		port:     port,
		database: database,
		user:     user,
		password: password,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	url := fmt.Sprintf("postgresql://%s:%s@%s:%v/%s", s.user, s.password, s.host, s.port, s.password)
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return err
	}
	s.conn = conn
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	err := s.conn.Close(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) CreateEvent(event storage.Event) (storage.Event, error) {
	return storage.Event{}, nil
}

func (s *Storage) UpdateEvent(id string, event storage.Event) (storage.Event, error) {
	return storage.Event{}, nil
}

func (s *Storage) DeleteEvent(id string) error {
	return nil
}

func (s *Storage) GetEventsDay(day time.Time) []storage.Event {
	return nil
}

func (s *Storage) GetEventsWeek(day time.Time) []storage.Event {
	return nil
}

func (s *Storage) GetEventsMonth(day time.Time) []storage.Event {
	return nil
}
