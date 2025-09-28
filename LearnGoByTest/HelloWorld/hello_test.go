package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Eva", "French")
		want := "Bonjour, Eva"
		assertCorrectMessage(t, got, want)
	})

}
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // tell the test suite this method is a helper, then the line number reported will
	// be in our func call rather than inside test helper.
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
