package context

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	data := make(chan string, 1)

	// We are simulating a slow process where we build the result slowly by appending the string, character by
	// character in a goroutine. When the goroutine finishes its work it writes the string to the data
	// channel.
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

type SpyResponseWriter struct {
	written bool
	t       *testing.T
}

func (s *SpyResponseWriter) Write(p []byte) (n int, err error) {
	s.written = true

	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true

	return nil
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyResponseWriter) assertWasWritten() {
	s.t.Helper()
	if !s.written {
		s.t.Error("response was not written")
	}
}

func (s *SpyResponseWriter) assertWasNotWritten() {
	s.t.Helper()
	if s.written {
		s.t.Error("response was written")
	}
}

func TestServer(t *testing.T) {
	data := "hello, world"

	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
		}

		// store.assertWasNotCancelled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{t: t}

		svr.ServeHTTP(response, request)

		response.assertWasNotWritten()
	})
}
