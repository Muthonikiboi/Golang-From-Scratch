package main

import "fmt"

func main(){
	fam:=map[string]float32{"joy":21,"lilian":17,"Dave":8}

	//access the values of map in golang
	fmt.Println(fam["joy"])

	//change the value of map
	fam["Dave"]=7
	fmt.Println(fam)

	//add an element to the map
	fam["Manu"]=22
	fmt.Println(fam)

	//delete an element in go
	delete(fam, "Dave")
	fmt.Println(fam)

	//looping through a map in Go
	for name,age := range fam {
		fmt.Printf("%s has %g years \n", name,age)
	}
}