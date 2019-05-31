package middleware

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type ContextKey string

const loggerContextKey ContextKey = "logger"

func RegisterLogger(next http.Handler, l *zap.SugaredLogger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		r.WithContext(context.WithValue(r.Context(), loggerContextKey, l))
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func Log(next http.Handler, l *zap.SugaredLogger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rww := NewResponseWriterWrapper(w)
		start := time.Now()
		defer func() {
			l.Infow(
				"Handling request",
				"bytes", rww.Bytes(),
				"status", rww.Status(),
				"time", time.Since(start),
			)
		}()
		next.ServeHTTP(rww, r)
	}
	return http.HandlerFunc(fn)
}

type responseWriterWrapper struct {
	http.ResponseWriter
	bytes         int
	status        int
	headerWritten bool
}

type ResponseWriterWrapper interface {
	http.ResponseWriter
	Bytes() int
	Status() int
}

func NewResponseWriterWrapper(w http.ResponseWriter) ResponseWriterWrapper {
	rww := &responseWriterWrapper{
		ResponseWriter: w,
		bytes:          0,
		status:         200,
		headerWritten:  false,
	}
	return rww
}

func (rww *responseWriterWrapper) Write(buf []byte) (int, error) {
	rww.WriteHeader(http.StatusOK)
	n, err := rww.ResponseWriter.Write(buf)
	rww.bytes += n
	return n, err
}

func (rww *responseWriterWrapper) WriteHeader(statusCode int) {
	if !rww.headerWritten {
		rww.status = statusCode
		rww.headerWritten = true
		rww.ResponseWriter.WriteHeader(statusCode)
	}
}

func (rww *responseWriterWrapper) Status() int {
	return rww.status
}

func (rww *responseWriterWrapper) Bytes() int {
	return rww.bytes
}
