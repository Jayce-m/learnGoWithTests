package helloworld

import "testing"

// This function must start with the word test
// it must take a single argument t *testing.T

func TestHello(t *testing.T) {
	t.Run("saying hello to peeps", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hello, Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("default behavior when empty string passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greeting in spanish", func(t *testing.T) {
		got := Hello("Jimmy", "Spanish")
		want := "Hola, Jimmy"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greeting in French", func(t *testing.T) {
		got := Hello("Ted", "French")
		want := "Bonjour, Ted"

		assertCorrectMessage(t, got, want)
	})

	t.Run("greeting in Italian", func(t *testing.T) {
		got := Hello("Marco", "Italian")
		want := "Ciao, Marco"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
}