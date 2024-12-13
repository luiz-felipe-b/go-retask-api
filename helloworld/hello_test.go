package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Luiz", "")
		want := "Hello, Luiz!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, world!' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Luiz", "Spanish")
		want := "Hola, Luiz!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Luiz", "French")
		want := "Bonjour, Luiz!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Portuguese", func(t *testing.T) {
		got := Hello("Luiz", "Portuguese")
		want := "Oi, Luiz!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
