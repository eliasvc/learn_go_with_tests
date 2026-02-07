package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("word in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q, but want %q given %q", got, want, "test")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "could not find the word you were looking for"

		if err == nil {
			t.Fatal("Expected error but didn't get one")
		}

		if err.Error() != want {
			t.Errorf("expected error %q, but got %q", want, err.Error())
		}

	})

}
