package main
import "fmt"

//without parameters
func addNumbers(){
 num1:=50
 num2:= 43

 sum:= num1+num2

 fmt.Println("Sum is: ", sum)
}


//with parameters and return value
func calcArea(length int,width int)int{

	area :=length*width
	

	fmt.Println("Area is:", area)
	return area
}

func main(){
	addNumbers()
	calcArea(40,21)
}