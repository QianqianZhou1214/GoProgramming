package arrays

func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers { //ignore the index value by using blank identifier
		sum += number
	}
	return sum
}
