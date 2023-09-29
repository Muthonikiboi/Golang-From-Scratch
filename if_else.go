package main
import "fmt"

func main() {
	number1 :=45
	number2 :=90

//if statement
    if number1==45{
		fmt.Printf("the value of number1 is %d \n", number1)
	}
//if else statement
    if number1<=number2{
		fmt.Printf("number2 is greater : %d \n", number2)
	}else{
		fmt.Printf("number1 is greater : %d \n", number1)
	}
//if...if else statement
    if number1==45 && number2==45{
		if number1==45{
			fmt.Printf("Both values are equal to 45: %d  & %d \n", number1 ,number2)
		}else{
			fmt.Printf("Both values are not equal to 45: %d & %d \n", number1 , number2)
		}
	}else{
		fmt.Printf("Invalid \n")
	}
//nested if statement
 //performance
 results :=87

 if results>91 && results<=100{
	fmt.Print("A \n")
 }else if results>81 && results<=90{
	fmt.Print("B \n")
 }else if results>71 && results<=80{
	fmt.Print("C \n")
 }else if results>61 && results<=70{
	fmt.Print("D \n")
 }else{
	fmt.Print("Failed \n")
 }
}