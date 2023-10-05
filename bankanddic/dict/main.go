package main

import (
	"fmt"
	"github.com/firstgo/bankanddic/dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "first word", "second": "hi"}
	fmt.Println(dictionary)
	err := dictionary.Add("first", "bb")
	dictionary.Update("first", "bbb")
	if err != nil {
		fmt.Println(err)
	}
	dictionary.Delete("first")
	fmt.Println(dictionary)
}
