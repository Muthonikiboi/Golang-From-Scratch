//Go print statement 
//string - %s
//integer - %d
//boolean - %t
//float - %g

package main
import "fmt"

func main(){
	//string and integer
	var name = "Joy"
	age := 23
	fmt.Printf("%s is a %d year old Software Engineer", name, age)

	//float
	fmt.Println("")
	var fuelPrice float32= 45.098765432
	fmt.Printf("The float number is : %g", fuelPrice)

	//boolean
	fmt.Println("")
	var boolean1 bool=false
	fmt.Printf("The boolean Value is :%t", boolean1)
}