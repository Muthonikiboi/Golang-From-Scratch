package main
import "fmt"

func main() {
	number1 := 34
	number2 :=56
	var result bool

	result=(number1==number2)
	fmt.Printf("the result is %t \n", result)

	result=(number2>number1)
	fmt.Printf("the result is %t", result)
}