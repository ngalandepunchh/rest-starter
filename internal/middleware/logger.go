package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Log(next http.Handler, l *zap.SugaredLogger) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		rww := NewResponseWriterWrapper(w)
		start := time.Now()
		defer func() {
			duration := time.Since(start)
			l.Infow(
				"Handling request",
				"bytesIn", r.ContentLength,
				"bytesOut", rww.Bytes(),
				"host", r.Host,
				"remote", r.RemoteAddr,
				"status", rww.Status(),
				"duration", uint64(duration/1000),
				"durationHuman", duration,
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
