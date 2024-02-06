package main

import (
	"fmt"
	"sort"
)

func main() {
	var numbers []int

	// ask for the size of the array
	var size int
	fmt.Println("Enter the size of array: ")
	fmt.Scan(&size)

	if size <= 0 {
		fmt.Println("The array is empty")
		return
	}

	// get input
	fmt.Println("Enter the number/s: ")

	numbers = make([]int, size)
	for i := range numbers {
		fmt.Scan(&numbers[i])
	}

	// sort the numbers
	sort.Ints(numbers)

	minimum := numbers[0]
	middle := (numbers[(len(numbers) / 2)])
	maximum := numbers[len(numbers)-1]

	if len(numbers)%2 == 0 {
		fmt.Println("The minimum is: ", minimum)
		fmt.Println("The maximum is: ", maximum)
	} else {
		fmt.Println("The minimum is: ", minimum)
		fmt.Println("The middle number is: ", middle)
		fmt.Println("The maximum is: ", maximum)
	}
}

/*package main

import (
	"fmt"
	"sort"
)

func main(){
	var a, b ,c int

	fmt.Println("Enter three number: ")
	fmt.Scan(&a, &b, &c)

	numbers := []int{a, b, c}
	sort.Ints(numbers)

	minumum := numbers[0]
	middle := numbers[1]
	maximum := numbers[2]

	fmt.Println("The minimum is: ", minumum)
	fmt.Println("The middle number is: ", middle)
	fmt.Println("The maximum is: ", maximum)
}*/
