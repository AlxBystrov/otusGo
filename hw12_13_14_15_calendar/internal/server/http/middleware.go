package internalhttp

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc { //nolint:unused
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		addr, _, _ := net.SplitHostPort(r.RemoteAddr)
		// TODO: add real status code for logging
		fmt.Printf("%s [%s] %s %s %s 200 %v \"%s\"\n", addr, time.Now().Format("02/Jan/2006:15:04:05 -0700"), r.Method, r.URL.Path, r.Proto, r.ContentLength, r.UserAgent())
	}
}
