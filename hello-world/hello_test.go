package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello with a name supplied", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello without a name supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello with a name supplied in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello with a name supplied in French", func(t *testing.T) {
		got := Hello("Louis", "French")
		want := "Bonjour, Louis!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
