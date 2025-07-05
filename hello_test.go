package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello with non-empty string", func(t *testing.T) {
		name := "Mushi"
		got := Hello(name)
		want := "Hello, " + name

		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello with empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
