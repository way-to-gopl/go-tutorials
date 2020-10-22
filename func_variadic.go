package main

import (
	"fmt"
)

//a variadic function is a function of indefinite arity, i.e., 
//one which accepts a variable number of arguments. More info:
//https://en.wikipedia.org/wiki/Variadic_function
func sum(numbers ...int) int {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	//use it like this:
	fmt.Println(sum(2, 3, 5, 7, 11, 13, 17, 19, 23, 29))

	//or like this:
	numbers := []int{31, 37, 41, 43, 47, 53, 59}
	fmt.Println(sum(numbers...))
}
