//Go takes Scan() function as the input
package main
import "fmt"

func main(){
	//There are three varients of Scan()
	//Scan
	var name string
	var age int

	fmt.Print("Enter your First name:")
	fmt.Scan(&name)
	fmt.Print("Enter age:")
	fmt.Scan(&age)
	fmt.Printf("Name: %s , Age: %d",name , age)
	fmt.Println("")

	//Scanln()
	var myName string

	fmt.Print("Enter Your last name :")
	fmt.Scanln(&myName)
	fmt.Printf("Last Name:%s",myName)
	fmt.Println("")

	//Scanf()
	var boolValue bool
	var float1 float64

	fmt.Print("Enter a boolean value :")
	fmt.Scanf("%t",&boolValue)
	fmt.Printf("the boolean value is :%t",boolValue)
	fmt.Println("")

	fmt.Print("Enter a float value :")
	fmt.Scanf("%g",&float1)
	fmt.Printf("The float value is : %g",float1)
}