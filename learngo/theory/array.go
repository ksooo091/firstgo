package theory

import "fmt"

func Arrayex() {
	// [배열 크기]타입{값}
	arrayexeample := [3]int{1, 2, 3}
	fmt.Println(arrayexeample)
}
func Sliceex() {
	// 다이나믹한 크기를 가진 배열
	sliceexeample := []int{1, 2, 3}
	// append로 값 추가시 다시 선언해야함
	sliceexeample = append(sliceexeample, 4)
	fmt.Println(sliceexeample)

}
