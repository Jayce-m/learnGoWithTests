package main

import "testing"

// This function must start with the word test
// it must take a single argument t *testing.T

func TestHello(t *testing.T) {
	got := Hello("Bob")
	want := "Hello, Bob"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}