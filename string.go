package main

import "fmt"

func main() {
	var message1 ="hello world"
	message2 := "world"
	fmt.Println(message1)
	fmt.Println(message2)

	//accessing character in a string
	fmt.Printf("%c\n",message2[2])

	//find the length of a string
	stringlength := len(message1)
	fmt.Println(len(stringlength))
}