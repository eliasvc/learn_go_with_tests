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

		assertDefinition(t, got, want, "test")
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
	t.Run("non-existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "new"
		want := "new word"

		err := dictionary.Add(word, want)
		if err != nil {
			t.Fatalf("unexpected error %q", err)
		}

		got, err := dictionary.Search(word)
		assertNoError(t, err)
		assertDefinition(t, got, want, word)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		err := dictionary.Add("test", "new test definition")
		if err == nil {
			t.Fatal("expected error, but didn't get one")
		}

		if err != ErrWordExists {
			t.Fatalf("expected error %q, but got %q", ErrWordExists, err)
		}

		// check definition wasn't changed
		currentDefintion, err := dictionary.Search("test")
		assertNoError(t, err)
		assertDefinition(t, currentDefintion, "this is just a test", "test")
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	newDefinition := "new definition"

	dictionary := Dictionary{word: definition}
	dictionary.Update(word, newDefinition)

	got, err := dictionary.Search(word)
	assertNoError(t, err)
	assertDefinition(t, got, newDefinition, word)
}

func assertDefinition(t testing.TB, got, want, word string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, but want %q given %q", got, want, word)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("expected no error, but got %q", got)
	}
}
