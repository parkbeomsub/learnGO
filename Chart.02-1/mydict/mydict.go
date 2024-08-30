package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not found")
	errWordExists = errors.New("Word Exists")
	errCantUpdate = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}

	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {

	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, definiton string) error {

	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definiton
	case errNotFound:
		return errCantUpdate
	}
	return nil

}

func (d Dictionary) Delete(word string) {
	delete(d, word)

}
