package theory

import "fmt"

func Pointerex() {
	a := 2
	// a의 메모리 위치
	b := &a
	// a값이 출력됨
	fmt.Println(*b)
	// a 값이 변경됨
	*b = 20

}
