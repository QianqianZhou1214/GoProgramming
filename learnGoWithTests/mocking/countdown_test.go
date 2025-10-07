package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	// backtick is another way to create String but lets you include things like new lines.

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
