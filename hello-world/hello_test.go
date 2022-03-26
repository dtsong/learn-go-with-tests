package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to someone", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("Say 'Hello World!' when an empty string is passed", func(t *testing.T) {
		got := Hello("")
		want := "Hello World!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
