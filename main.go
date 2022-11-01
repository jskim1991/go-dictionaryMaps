package main

import (
	"dictionary/dictionary"
	"fmt"
)

func main() {
	dictionary := dictionary.Dictionary{"first": "First word"}
	dictionary.Add("second", "Second word")

	addError := dictionary.Add("second", "Second word")
	if addError != nil {
		fmt.Println(addError)
	}

	definition, error := dictionary.Search("second")
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(definition)
	}

	dictionary.Update("first", "THE FIRST word")
	dictionary.Delete("second")
	fmt.Println(dictionary)
}
