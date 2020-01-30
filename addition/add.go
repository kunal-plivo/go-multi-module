package addition

import "fmt"

func addNumbers(numbers ...int) int {
	fmt.Printf("Numbers = %s", numbers)
	sum := 0
	for _, num := range numbers {
		sum = sum + num
	}
	return sum
}