package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//Table driven tests
	areaTests := []struct { //slice of struct
		shape Shape   //field
		want  float64 //field
	}{
		// fill slice with cases
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper() //our helper does not need to concern whether it's a circle or a rectangle.
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkArea(t, rectangle, rectangle.Area())
	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, circle.Area())
	})
	/*
		t.Run("rectangle", func(t *testing.T) {
			rectangle := Rectangle{12, 6}
			got := rectangle.Area()
			want := 72.0
			if got != want {
				t.Errorf("got %.2f want %.2f", got, want)
			}
		})

		t.Run("circle", func(t *testing.T) {
			circle := Circle{10}
			got := circle.Area()
			want := 314.1592653589793
			if got != want {
				t.Errorf("got %g want %g", got, want) //g: more precise
			}
		}

	*/

}
