package iteration

func repeat(char string, iterationQuantity int) string {
	var result string

	for i := 0; i < iterationQuantity; i++ {
		result += char
	}

	return result
}
