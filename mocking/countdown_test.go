package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)
	want := "3"
	got := buffer.String()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
