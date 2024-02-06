package logger

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {

	t.Run("basic logger test", func(t *testing.T) {
		var buf bytes.Buffer
		logg := New(&buf, "debug", "text")

		logg.Debug("Debug message")
		currentTime := time.Now().Format("2006-01-02T15:04:05.000Z07:00")
		expectedMessage := "time=" + currentTime + " level=DEBUG msg=\"Debug message\"\n"
		require.Equal(t, expectedMessage, buf.String())

		buf.Reset()
		logg.Info("Info message")
		currentTime = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
		expectedMessage = "time=" + currentTime + " level=INFO msg=\"Info message\"\n"
		require.Equal(t, expectedMessage, buf.String())

		buf.Reset()
		logg.Warning("Warning message")
		currentTime = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
		expectedMessage = "time=" + currentTime + " level=WARN msg=\"Warning message\"\n"
		require.Equal(t, expectedMessage, buf.String())

		buf.Reset()
		logg.Error("Error message")
		currentTime = time.Now().Format("2006-01-02T15:04:05.000Z07:00")
		expectedMessage = "time=" + currentTime + " level=ERROR msg=\"Error message\"\n"
		require.Equal(t, expectedMessage, buf.String())

		logg = New(&buf, "warning", "text")
		buf.Reset()
		logg.Debug("some text, that never be printed")
		expectedMessage = ""
		require.Equal(t, expectedMessage, buf.String())
	})
}
