package arraysandslices

func sum(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result +=
			number
	}
	return result
}

func sumAllRest(numbers ...[]int) []int {
	var result []int

	for _, number := range numbers {
		if len(number) == 0 {
			result = append(result, 0)
		} else {

			final := number[1:]
			result = append(result, sum(final))
		}
	}

	return result
}
