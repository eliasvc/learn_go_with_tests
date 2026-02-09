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

		if err == nil {
			t.Fatal("Expected error but didn't get one")
		}

		if err != ErrNotFound {
			t.Errorf("expected error %q, but got %q", ErrNotFound, err)
		}
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "new"
	want := "new word"

	dictionary.Add(word, want)

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("expected no error, but got %q", err)
	}

	if want != got {
		t.Errorf("got %q, but want %q given %q", got, want, word)
	}

}
