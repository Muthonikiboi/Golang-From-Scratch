package main

import (
	"fmt"
	"strings"
)

func main() {
	var message1 ="hello world"
	message2 := "world"
	fmt.Println(message1)
	fmt.Println(message2)

	//accessing character in a string
	fmt.Printf("%c\n",message2[2])

	//find the length of a string
	//stringlength := len(message1)
	//fmt.Println(len(stringlength))

	// to join two strings
	message3:= "I love"
	message4 :="go programming language"

	result := message3 + " " + message4
	fmt.Println(result)

	//GO STRING FUNCTIONS
	//COMPARE()
	//CONTAINS()
	//REPLACES()
	//ToLower()
	//ToUpper()
	//Split()

	//compare two stings
	string1 :="Programiz"
	string2 :="Programi"
	string3 :="Programiz"

	fmt.Println(strings.Compare(string1,string2))
	fmt.Println(strings.Compare(string1,string3))
	fmt.Println(strings.Compare(string2,string3))
}