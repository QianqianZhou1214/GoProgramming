package maps

import "errors"

type Dictionary map[string]string

// do not do var m map[string]string to initialize a nil map variable
// do make map: var dictionary = map[string]string{} OR var dictionary = make(map[string]string)

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
	// A map value is a pointer to a runtime.hmap structure.
}
