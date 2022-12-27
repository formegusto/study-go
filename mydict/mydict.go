package mydict

import (
	"fmt"

	"github.com/formegusto/study-go/myerrors"
)

// 단순히 type에 대한 alias를 걸어놓은 것 이다.
type Money int 

func MoneyTest() {
	fmt.Println(Money(1))
	// -> 1
}

// Dictionary type
type Dictionary map[string]string 

// Search for a word
func (d Dictionary) Search(word string) (string,error) {
	// map의 독특한 기능
	// 두 번째에는 해당 키의 값이 존재하는지 찾아주는 것 이다.
	value, exists := d[word]

	if exists {
		return value, nil
	}
	return "",myerrors.NotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error{
	// 있으면 추가를 안한다는 내용
	_, err := d.Search(word)
	switch err {
		case myerrors.NotFound:
			d[word] = def
		case nil:
			return myerrors.WordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
		case nil:
			d[word] = def
		case myerrors.NotFound:
			return myerrors.CantUpdate
	}

	return nil
}

func (d Dictionary) Delete(word string) error{
	_, err := d.Search(word)
	switch err {
		case nil:
			delete(d, word)
		case myerrors.NotFound:
			return myerrors.CantDelete
	}
	return nil
}