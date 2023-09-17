package main
import "fmt"

func main() {
//how to declare variables
var number int =23
fmt.Println(number)

var name string = "Joy"
fmt.Println(name)

//different methods to declare variables
//METHOD1
var name1 string = "Getange"
fmt.Println(name1)
//METHOD2-we do not declare the datatype
var name3 ="Lily"
fmt.Println(name3)
//METHOD3
name4 :="Hello"
fmt.Println(name4)

//declaring many variables in one line
var myName,contact,school,location= "Joy",9888888888,"Dekut","Nyeri"
fmt.Println(myName,contact,school,location)

//Changing the value of a variable
var year int=2022
fmt.Println(year)

year=2023
fmt.Println(year)//output of year willbe 2023

//to declaire constants in go
//const speed=2000
//speed=4567
//fmt.Println(speed)
//cannot change value once declaired
}