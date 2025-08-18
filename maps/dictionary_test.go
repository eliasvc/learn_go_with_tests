package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word exists in dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("unknown")
		want := ErrNotFound

		if err == nil {
			t.Fatal("expected error but didn't get one")
		}

		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("add", "testing adding a word")

		got, err := dictionary.Search("add")
		want := "testing adding a word"

		if err != nil {
			t.Fatal("unexpected error adding word:", err)
		}

		assertStrings(t, got, want)
	})

	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})
}

func TestUpdate(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("unknown word", func(t *testing.T) {
		word := "test"
		newDefinition := "new definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)
	})

}

func TestDelete(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		assertError(t, err, nil)

		_, err = dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Delete("test")
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("unexpected error searching for word %q: %s", word, err)
	}

	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
