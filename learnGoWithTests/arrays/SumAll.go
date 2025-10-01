package arrays

func SumAll(numbersToSum ...[]int) []int { // this is a Variadic function
	// the type of numbersToSum is equivalent to []int,
	//which means we can use len() and iterates over it with range.
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

// fmt.Println() is a common variadic function (with any number of trailing arguments)
