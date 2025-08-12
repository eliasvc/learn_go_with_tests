package main

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numbersToSum {
		sum = append(sum, Sum(numbers))
	}

	return sum
}
