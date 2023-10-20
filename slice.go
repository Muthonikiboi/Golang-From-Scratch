package main

import "fmt"

func main() {
	numbers:=[] int {2,56,78,90,34,56}
	fmt.Println(numbers)

	//create a slice from a go array
	number1 := [7] int {2,4,5,8,90,12,4}
	sliceNumbers := number1[3:7]
	fmt.Println(sliceNumbers)

	//copy
	primeNumbers:=[]int{2,3,5,7}
	number2:=[]int{1,2,3}
	copy(number2, primeNumbers)
	fmt.Println(number2)//the last prime number is not chosen since the number2 array has a value of 3

	//append--adding to the array
	primeNumbers= append(primeNumbers,11,17)

	fmt.Println(primeNumbers)
	
	//looping go slices
	for i:=0; i<len(numbers);i++{
		fmt.Println(numbers[i])
	}
}
