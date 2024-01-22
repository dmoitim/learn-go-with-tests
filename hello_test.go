package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("DevA")
	want := "Hello, DevA"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
