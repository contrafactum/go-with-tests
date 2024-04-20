package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Danny")
	want := "Hello Danny"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
