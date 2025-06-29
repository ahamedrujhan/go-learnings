package main

import "fmt"

type transformerFunc func(int) int

func createTransformer(fact int) transformerFunc {
	return func(number int) int {
		return number * fact
	}
}

func transformed(numbers *[]int, transformer transformerFunc) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transformer(val))
	}
	return dNumbers
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	doubleTransformer := createTransformer(2)

	tripleTransformer := createTransformer(3)

	double := transformed(&numbers, doubleTransformer)
	triple := transformed(&numbers, tripleTransformer)

	fmt.Println("double numbers :- ", double)
	fmt.Println("triple numbers :- ", triple)

}
