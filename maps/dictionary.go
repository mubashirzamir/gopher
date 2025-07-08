package maps

import "fmt"

// you should never initialize a nil map variable:
// var m map[string]string

// Instead, you can initialize an empty map or use the make keyword to create a map for you:
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return "DictionaryErr: " + string(e)
}

func ErrWordNotFound(word string) DictionaryErr {
	return DictionaryErr(fmt.Sprintf("could not find %q", word))
}

func ErrWordAlreadyExists(word, definition string) DictionaryErr {
	return DictionaryErr(fmt.Sprintf("%q already exists with definition %q", word, definition))
}

func ErrWordDoesNotExist(word string) DictionaryErr {
	return DictionaryErr(fmt.Sprintf("%q does not exist", word))
}

func ErrWordNotFoundForDelete(word string) DictionaryErr {
	return DictionaryErr(fmt.Sprintf("%q already does not exist", word))
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound(word)
	}

	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	existingDefinition, err := d.Search(word)

	switch err {
	case ErrWordNotFound(word):
		d[word] = definition
	case nil:
		return ErrWordAlreadyExists(word, existingDefinition)
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound(word):
		return ErrWordDoesNotExist(word)
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound(word):
		return ErrWordNotFoundForDelete(word)
	case nil:
		delete(d, word)
	}

	return nil
}
