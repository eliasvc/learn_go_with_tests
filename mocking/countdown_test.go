package main

import (
	"bytes"
	"reflect"
	"testing"
)

const sleep = "sleep"
const write = "write"

type SleepCountdownOperations struct {
	Calls []string
}

func (s *SleepCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SleepCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SleepCountdownOperations{})

		want := `3
2
1
Go!`
		got := buffer.String()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrint := &SleepCountdownOperations{}
		Countdown(spySleeperPrint, spySleeperPrint)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleeperPrint.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperPrint.Calls)
		}
	})
}
