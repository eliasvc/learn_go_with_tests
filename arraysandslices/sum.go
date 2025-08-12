package main

func Sum(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sum []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
			continue
		}
		sum = append(sum, Sum(numbers[1:]))
	}

	return sum
}
