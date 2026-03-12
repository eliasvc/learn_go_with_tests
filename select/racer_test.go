package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("faster server response wins", func(t *testing.T) {

		slowServer := makeDelayedSever(20 * time.Millisecond)
		fastServer := makeDelayedSever(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		fastURL := fastServer.URL
		slowURL := slowServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("didn't expect error, but got %q", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns error if server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayedSever(25 * time.Millisecond)

		defer slowServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, slowServer.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedSever(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
	}))
}
