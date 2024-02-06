package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/AlxBystrov/otusGo/hw12_13_14_15_calendar/internal/storage"
)

type Storage struct {
	mu     sync.RWMutex //nolint:unused
	events map[string]storage.Event
}

func New() *Storage {
	return &Storage{
		events: make(map[string]storage.Event),
	}
}

var (
	ErrEventIDExists  = errors.New("the specified ID is already in use")
	ErrEventIDMissing = errors.New("the specified ID is not present")
)

func (s *Storage) CreateEvent(event storage.Event) (storage.Event, error) {
	_, ok := s.events[event.ID]
	if ok {
		return storage.Event{}, ErrEventIDExists
	}
	s.mu.Lock()
	s.events[event.ID] = event
	s.mu.Unlock()
	return s.events[event.ID], nil
}

func (s *Storage) UpdateEvent(id string, event storage.Event) (storage.Event, error) {
	_, ok := s.events[event.ID]
	if !ok {
		return storage.Event{}, ErrEventIDMissing
	}
	s.mu.Lock()
	s.events[id] = event
	s.mu.Unlock()
	return s.events[id], nil
}

func (s *Storage) DeleteEvent(id string) error {
	_, ok := s.events[id]
	if !ok {
		return ErrEventIDMissing
	}
	s.mu.Lock()
	delete(s.events, id)
	s.mu.Unlock()
	return nil
}

func (s *Storage) GetEventsDay(day time.Time) []storage.Event {
	var result []storage.Event

	for _, event := range s.events {
		if event.Date.Equal(day) {
			result = append(result, event)
		}
	}
	return result
}

func (s *Storage) GetEventsWeek(day time.Time) []storage.Event {
	var result []storage.Event

	for _, event := range s.events {
		if event.Date.After(day) && event.Date.Before(day.Add(time.Hour*24*7)) {
			result = append(result, event)
		}
	}
	return result
}

func (s *Storage) GetEventsMonth(day time.Time) []storage.Event {
	var result []storage.Event

	for _, event := range s.events {
		if event.Date.After(day) && event.Date.Before(day.Add(time.Hour*24*30)) {
			result = append(result, event)
		}
	}
	return result
}
