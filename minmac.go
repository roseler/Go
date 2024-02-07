package main

import (
	"fmt"
	"sort"
)

func main() {package main

import "fmt"

func main() {
	var myNum []int

	var size int
	fmt.Println("Enter the size of array: ")
	fmt.Scan(&size)

	if size <= 0 {
		fmt.Println("The array is empty")
		return
	}

	fmt.Println("Enter the number/s: ")

	myNum = make([]int, size)
	for i := range myNum {
		fmt.Scan(&myNum[i])
	}

	// sort the myNum
	for i := 0; i < len(myNum); i++ {
		for j := 0; j < len(myNum)-1; j++ {
			if myNum[j] > myNum[j+1] {
				myNum[j], myNum[j+1] = myNum[j+1], myNum[j]
			}
		}
	}

	minimum := myNum[0]
	middle := (myNum[(len(myNum) / 2)])
	maximum := myNum[len(myNum)-1]

	if len(myNum)%2 == 0 {
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

func main() {
	var myNum []int

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

	myNum = make([]int, size)
	for i := range myNum {
		fmt.Scan(&myNum[i])
	}

	// sort the myNum
	sort.Ints(myNum)

	minimum := myNum[0]
	middle := (myNum[(len(myNum) / 2)])
	maximum := myNum[len(myNum)-1]

	if len(myNum)%2 == 0 {
		fmt.Println("The minimum is: ", minimum)
		fmt.Println("The maximum is: ", maximum)
	} else {
		fmt.Println("The minimum is: ", minimum)
		fmt.Println("The middle number is: ", middle)
		fmt.Println("The maximum is: ", maximum)
	}
}
*/
/*package main

import (
	"fmt"
	"sort"
)

func main(){
	var a, b ,c int

	fmt.Println("Enter three number: ")
	fmt.Scan(&a, &b, &c)

	myNum := []int{a, b, c}
	sort.Ints(myNum)

	minumum := myNum[0]
	middle := myNum[1]
	maximum := myNum[2]

	fmt.Println("The minimum is: ", minumum)
	fmt.Println("The middle number is: ", middle)
	fmt.Println("The maximum is: ", maximum)
}*/

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
