package mapCh

import "testing"

func TestSearch(t *testing.T) {

	d := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := d.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := d.Search("unknown")
		if got == nil {
			t.Fatal("expected an error")
		}

		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := d.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, d, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		d := Dictionary{word: definition}

		err := d.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{word: definition}
		updatedDef := "update test"

		err := d.Update(word, updatedDef)
		assertError(t, err, nil)
		assertDefinition(t, d, word, updatedDef)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		d := Dictionary{}

		err := d.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		d := Dictionary{word: "test definition"}

		err := d.Delete(word)
		assertError(t, err, nil)

		_, err = d.Search(word)
		assertError(t, err, ErrorNotFound)
	})

	t.Run("non-existing-word", func(t *testing.T) {
		word := "test"
		d := Dictionary{}

		err := d.Delete(word)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertDefinition(t testing.TB, dict Dictionary, word string, definition string) {
	t.Helper()

	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find added word")
	}
	assertString(t, got, definition)
}
