package arrays

func SumAll(numbersToSum ...[]int) []int { // this is a Variadic function
	// the type of numbersToSum is equivalent to []int,
	//which means we can use len() and iterates over it with range.
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) //this returns a new slice to avoid runtime error when you are out of bound.
	}
	return sums
}

// fmt.Println() is a common variadic function (with any number of trailing arguments)

// make([]int, 0, 5): allows to create a slice with a starting capacity of the len(), which is the number of elements.
// capacity(5) is the number of elements it can hold in underlying array.
