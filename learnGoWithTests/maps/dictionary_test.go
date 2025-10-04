package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	// key type can only be comparable type; value type can be any even a map

	got := Search(dictionary, "test")
	want := "this is just a test"
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}
