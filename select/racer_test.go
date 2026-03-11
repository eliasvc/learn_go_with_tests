package main

import "testing"

func TestRacer(t *testing.T) {
	fastURL := "https://www.facebook.com"
	slowURL := "https://www.quii.dev"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
