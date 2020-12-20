package mydict

import "errors"

// Dictionary type
/* NOTE:
type Dictionary
- map에 대한 alias(가명)을 지정 (struct가 아님)
- type은 method를 가질 수 있음
*/
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("Already Exists")
	errCantUpdate = errors.New("Cannot Update")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	// map[key]의 return value는 (value, exist(존재여부 bool)임
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to dictionary
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

// Update dict
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
	// map의 delete 함수 : key가 존재하지 않으면 아무것도 실행하지 않음
}
