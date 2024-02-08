package sqlstorage

import (
	"context"
	"fmt"
	// "database/sql"
	"time"

	"github.com/AlxBystrov/otusGo/hw12_13_14_15_calendar/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	host     string
	port     int
	database string
	user     string
	password string
	conn     *pgxpool.Pool
	ctx      context.Context
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
	pgxpool, err := pgxpool.New(ctx, url)
	if err != nil {
		return err
	}
	s.conn = pgxpool
	s.ctx = ctx
	return nil
}

func (s *Storage) Close(ctx context.Context) {
	s.conn.Close()
}

func (s *Storage) CreateEvent(event storage.Event) (storage.Event, error) {
	newEvent := storage.Event{}
	q := `insert into events(id, title, ts, duration, description, user_id, notify_before) 
	      values($1, $2, $3, $4, $5, $6, $7)
		  returning id, title, ts, duration, description, user_id, notify_before`
	result, err := s.conn.Query(s.ctx, q, event.ID, event.Title, event.Date, event.Duration, event.Description, event.UserID, event.NotifyBefore)
	if err != nil {
		return storage.Event{}, err
	}
	err = result.Scan(&newEvent.ID, &newEvent.Title, &newEvent.Date, &newEvent.Duration, &newEvent.Description, &newEvent.UserID, &newEvent.NotifyBefore)
	if err != nil {
		return newEvent, err
	}
	return newEvent, nil
}

func (s *Storage) UpdateEvent(id string, event storage.Event) (storage.Event, error) {
	fixedEvent := storage.Event{}
	q := `UPDATE events 
	      SET title=$1, ts=$2, duration=$3, description=$4, user_id=$5, notify_before=$6
		  WHERE id=$8
		  RETURNING id, title, ts, duration, description, user_id, notify_before`
	result, err := s.conn.Query(s.ctx, q, event.Title, event.Date.Unix(), event.Duration, event.Description, event.UserID, event.NotifyBefore, event.ID)
	if err != nil {
		return storage.Event{}, err
	}
	err = result.Scan(&fixedEvent.ID, &fixedEvent.Title, &fixedEvent.Date, &fixedEvent.Duration, &fixedEvent.Description, &fixedEvent.UserID, &fixedEvent.NotifyBefore)
	if err != nil {
		return fixedEvent, err
	}
	return fixedEvent, nil
}

func (s *Storage) DeleteEvent(id string) error {
	q := `DELETE FROM events
	      WHERE id=$1`
	_, err := s.conn.Query(s.ctx, q, id)
	if err != nil {
		return err
	}
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
