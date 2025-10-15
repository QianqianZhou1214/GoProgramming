package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})
	// call func walk with a struct(x) has a string field in it.

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d, want %d", len(got), 1)
	}
}
