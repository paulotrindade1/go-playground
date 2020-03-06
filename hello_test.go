package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func (t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func (t *testing.T) {
		got := Hello("Paulo", "Portuguese")
		want := "Olá, Paulo"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func (t *testing.T) {
		got := Hello("Paulo", "French")
		want := "Bonjour, Paulo"

		assertCorrectMessage(t, got, want)
	})

}