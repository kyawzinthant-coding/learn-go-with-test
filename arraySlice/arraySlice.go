package arrayandslice

func Sum(numbers []int) int {
	result := 0

	for _, nunumber := range numbers {
		result += nunumber
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	numerOfLength := len(numbersToSum)
	sum := make([]int, numerOfLength)

	for i, number := range numbersToSum {
		sum[i] = Sum(number)
	}

	return sum
}
