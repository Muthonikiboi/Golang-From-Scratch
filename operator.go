package main
import "fmt"

func main(){
	//Arithmetic Operators
	//addition
	num1:= 79
	num2:=56

	sum:= num1 + num2
	fmt.Printf("%d + %d=%d\n", num1, num2 ,sum)

	//subtraction
	difference:=num1-num2
	fmt.Printf("%d -%d=%d\n", num1,num2, difference)

	//division-integer
	division:= num1/num2
	fmt.Printf("%d/ %d =%d\n", num1,num2, division)

	//modulous
	remainder:= num1%num2
	fmt.Println(remainder)

	//incrteamental and decreamental
	num1++
	fmt.Println(num1)

	num2--
	fmt.Println(num2)

	//assignment Operator
	num:=19
	var result int
	result=num
	fmt.Println(result)

	//logical operators
}