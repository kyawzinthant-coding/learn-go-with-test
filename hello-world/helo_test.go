package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kyaw")
	want := "Hello, Kyaw"

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
