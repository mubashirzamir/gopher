package select_demo

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const timeout = 10 * time.Second // seconds

func TestRacer(t *testing.T) {
	t.Run("returns fastest URL", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(timeout, slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestTimeout(t *testing.T) {
	t.Run("returns an error if a server doesn't respond within defined timeout", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(timeout, serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected error but did not get one.")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
