package main
import "fmt"

func main(){
	var sum = func(a , b int){

		sum:=a+b

		fmt.Println(sum)
	}

	sum(23,90)

	//return value from anonymous function 
	area := func(length,width int) int {
		return length * width
	}

	//function call using variable name 
	fmt.Println("The area of the rectangle is",area(3,4))
}
