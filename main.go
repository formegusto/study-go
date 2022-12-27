package main

import (
	"fmt"

	"github.com/formegusto/study-go/mydict"
)

func SearchAndPrint(dictionary mydict.Dictionary, searchWord string) {
	def, err := dictionary.Search(searchWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(def)
	}
}

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}

	searchWord := "second"
	SearchAndPrint(dictionary, searchWord)

	word := "second"
	// addErr := dictionary.Add("first", "Second word")
	addErr := dictionary.Add(word, "Second word")
	if addErr != nil {
		fmt.Println(addErr)
	}
	SearchAndPrint(dictionary, word)

	// badWord := "asfadsfas"
	updateErr := dictionary.Update(word, "Update Second Word")
	if updateErr != nil {
		fmt.Println(updateErr)
	}
	SearchAndPrint(dictionary, word)

	// deleteErr := dictionary.Delete(badWord)
	deleteErr := dictionary.Delete(word)
	if deleteErr != nil {
		fmt.Println(deleteErr)
	}
	SearchAndPrint(dictionary, word)


	// fmt.Println(dictionary["first"])
	// dictionary["hello"] = "hello"
	// fmt.Println(dictionary)
}