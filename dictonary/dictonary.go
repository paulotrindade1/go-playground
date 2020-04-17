package dictonary

import "errors"

type Dictonary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictonary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictonary) Add(word, definition string) {
	d[word] = definition
}
