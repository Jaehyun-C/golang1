package main

import (
	"fmt"

	"github.com/Jaehyun-C/golang1/learn2/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")

	definition, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	dictionary.Update(baseWord, "Second")

	definition, err = dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}

	dictionary.Delete(baseWord)
}
