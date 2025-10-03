package structs

import (
	"testing"
)

type Rectangle struct { // A struct is a named collection of fields where you can store data.
	Width  float64
	Height float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{12, 6}
	got := Area(rectangle)
	want := 72.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
