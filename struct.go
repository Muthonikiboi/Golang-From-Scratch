package main
import "fmt"

func main(){
	type Person struct{
		name string
		course string 
		age int
	}

person1 :=Person{"Joy","IT",21}
fmt.Println(person1)


var person2 Person
person2=Person{"Manu","IT",22}
fmt.Println(person2)

//accessing a struct
fmt.Println("Pewrson1's Name:",person1.name)
}