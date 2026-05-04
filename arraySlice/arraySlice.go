package arrayandslice

func Sum(numbers []int) int {
	result := 0

	for _, nunumber := range numbers {
		result += nunumber
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int

	for _, number := range numbersToSum {
		sum = append(sum, Sum(number))
	}

	return sum
}
