package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string

var (
	errorNotFound   = errors.New("찾을수 없음")
	errorWordExists = errors.New("이미 존재하는 키")
	errorCantUpdate = errors.New("키가 존재 하지 않아 업데이트 불가")
	errorCantDelete = errors.New("키가 존재 하지 않아 삭제 불가")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errorNotFound
}

// Add a word to the dic
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch {
	case errors.Is(err, errorNotFound):
		d[key] = value
	case err == nil:
		return errorWordExists
	}

	return nil
}

// Update a value to the dic
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch {
	case errors.Is(err, errorNotFound):
		return errorCantUpdate
	case err == nil:
		d[key] = value
	}
	return nil
}

// Delete key
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch {
	case errors.Is(err, errorNotFound):
		return errorCantDelete
	case err == nil:
		delete(d, key)
	}
	// 키 없으면 암것도 안하고 있을때만 삭제할수 있긴함.

	return nil
}
