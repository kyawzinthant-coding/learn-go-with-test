package arrayandslice

func Sum(numbers []int) int {
	result := 0

	for _, nunumber := range numbers {
		result += nunumber
	}
	return result
}
