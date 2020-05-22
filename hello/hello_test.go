package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf(
				"got %q want %q", got, want,
			)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Joe", "")
		want := "Hello, Joe!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say's 'Hello, world!' when empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Pierre", "French")
		want := "Bonjour, Pierre!"
		assertCorrectMessage(t, got, want)
	})
}
