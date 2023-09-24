package theory

func Foradd(numbers ...int) int {
	result := 0
	// for 값 하나만 주면 range값 뱉음. 따라서 index 넣어주거나 _ 넣어줘서 생략해야함.
	for _, number := range numbers {
		result += number
	}
	// 위 범위는 아래와 같다.
	for i := 0; i < len(numbers); i++ {
		// 변수 대신 numbers[i] 사용

	}

	return result
}
