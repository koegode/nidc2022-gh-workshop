package main

import (
	"testing"
)

func TestTauntMe(t *testing.T) {
	got := TauntMe("Glenn")
	want := "Glenn smells of elderberries."

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGiveMePi(t *testing.T) {
	got := GiveMePi()
	want := 3.14

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}
}

