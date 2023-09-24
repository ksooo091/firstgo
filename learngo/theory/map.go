package theory

import "fmt"

func Mapex() {
	// map [키 타입]밸류타입
	example := map[int]string{1: "a"}
	fmt.Println(example)
	for key, value := range example {
		fmt.Println(key, value)
	}
}

type exampleType struct {
	key1 string
	key2 int
	key3 bool
	key4 []string
}

func Stuctsex() {
	key4 := []string{"a", "b", "c"}
	example := exampleType{key1: "a", key2: 1, key3: true, key4: key4}
	fmt.Println(example)

}
