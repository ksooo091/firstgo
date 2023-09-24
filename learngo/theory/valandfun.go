package theory

import (
	"fmt"
	"strings"
)

func cal(a, b int) int {
	return (a + b)

}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}
func lenAndUpper2(name string) (lenght int, uppername string) {
	defer fmt.Println("This message will printed when func done")
	lenght = len(name)
	uppername = strings.ToUpper(name)
	return
}

func valex() {
	// 상수
	const test1 string = "test1 const"
	// 변수 , 아래 두가지 선언방식은 같다.
	var test2 string = ""
	test3 := ""

	fmt.Println(test1)
	fmt.Println(test2)
	fmt.Println(test3)
	//멀티리턴으로 값 여러개 받기
	muiltReturnInt, muiltReturnString := lenAndUpper("myname")
	// 멀티 리턴으로 값 받을때 특정 값 생략하려면 _ 로 값 주면 됨.
	_, muiltReturn := lenAndUpper("asdfff")
	fmt.Println(muiltReturnInt, muiltReturnString)
	fmt.Println(muiltReturn)

}
