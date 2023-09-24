package theory

func Ifex(input int) string {
	// if문을 위한 변수 생성
	if example := input + 1; example > 10 {
		return "true"
	}
	return "false"
}

func Switchex(input int) string {
	// if문처럼 전용 변수 생성 가능
	switch example := input + 1; {
	case example < 3:
		return "true"
	case example > 5:
		return "x"
	}
	switch input {
	case 3:
		return "true"
	case 15:
		return "x"
	}

	return "false"
}
