package interation

const interationQuantity = 5

func repeat(char string) string {
	var result string

	for i := 0; i < interationQuantity; i++ {
		result += char
	}

	return result
}
