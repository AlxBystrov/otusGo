package memorystorage

import (
	"testing"
	"time"

	"github.com/AlxBystrov/otusGo/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {

	events := []storage.Event{
		{
			ID: "1", Title: "one",
			Date:     time.Now().Add(time.Hour),
			Duration: time.Hour, Description: "description 1",
			UserID: "1", NotifyBefore: time.Hour,
		},
		{
			ID: "2", Title: "two",
			Date:     time.Now().Add(time.Minute * 30),
			Duration: time.Hour, Description: "description 2",
			UserID: "1", NotifyBefore: time.Hour,
		},
		{
			ID: "3", Title: "three",
			Date:     time.Now().Add(time.Hour * 24 * 5),
			Duration: time.Hour, Description: "description 3",
			UserID: "2", NotifyBefore: time.Hour,
		},
		{
			ID: "4", Title: "four",
			Date:     time.Now().Add(time.Hour * 24 * 10),
			Duration: time.Hour, Description: "description 4",
			UserID: "2", NotifyBefore: time.Hour,
		},
	}

	storageTest := New()

	t.Run("add events test", func(t *testing.T) {
		for _, event := range events {
			storedEvent, err := storageTest.CreateEvent(event)
			require.Equal(t, nil, err)
			require.Equal(t, event, storedEvent)
		}
	})

	t.Run("get events day test", func(t *testing.T) {
		ev := storageTest.GetEventsDay(time.Now())
		require.Equal(t, 2, len(ev))
	})

	t.Run("get events week test", func(t *testing.T) {
		ev := storageTest.GetEventsWeek(time.Now())
		require.Equal(t, 3, len(ev))
	})

	t.Run("get events month test", func(t *testing.T) {
		ev := storageTest.GetEventsMonth(time.Now())
		require.Equal(t, 4, len(ev))
	})

	newEvent := storage.Event{
		ID: "3", Title: "three",
		Date:     time.Now().Add(time.Hour * 48),
		Duration: time.Hour, Description: "description 3",
		UserID: "2", NotifyBefore: time.Hour,
	}

	t.Run("update event test", func(t *testing.T) {
		updatedEvent, err := storageTest.UpdateEvent(newEvent.ID, newEvent)
		require.Equal(t, nil, err)
		require.Equal(t, newEvent, updatedEvent)
	})

	t.Run("delete event test", func(t *testing.T) {
		err := storageTest.DeleteEvent(newEvent.ID)
		require.Equal(t, nil, err)
		ev := storageTest.GetEventsMonth(time.Now())
		require.Equal(t, 3, len(ev))
	})
}
