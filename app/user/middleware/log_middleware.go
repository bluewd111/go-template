package middleware

import (
	"net/http"
	"time"
)

type LogModel struct {
	remoteAddr    string
	userName      string
	loggedAt      string
	httpMethod    string
	path          string
	proto         string
	statusCode    int
	contentLength int64

	scheme    string
	host      string
	userAgent string
	referer   string
}

func NewLogModel(r *http.Request) *LogModel {
	return &LogModel{
		remoteAddr:    r.RemoteAddr,
		userName:      r.URL.User.Username(),
		loggedAt:      time.Now().Format("2006-01-02T15:04:05+09:00"),
		httpMethod:    r.Method,
		path:          r.URL.Path,
		proto:         r.Proto,
		statusCode:    r.Response.StatusCode,
		contentLength: r.Response.ContentLength,

		scheme:    r.URL.Scheme,
		host:      r.URL.Host,
		userAgent: r.UserAgent(),
		referer:   r.Referer(),
	}
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// log.Println(NewLogModel(r))
		next.ServeHTTP(w, r)
	})
}
