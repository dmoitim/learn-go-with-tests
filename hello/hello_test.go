package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		got := Hello("Devair", "")
		want := "Hello, Devair"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Deva", "Spanish")
		want := "Hola, Deva"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Deva", "French")
		want := "Bonjour, Deva"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Portuguese", func(t *testing.T) {
		got := Hello("Deva", "Portuguese")
		want := "Olá, Deva"
		assertCorrectMessage(t, got, want)
	})
}
