//Range...

package main
import "fmt"

func main(){

	numbers := []int{78,56,34,12,34,78,90,34,56,78,13,45,80,65,73}

	for index,numbers := range numbers{
		fmt.Printf("The value id index %d with value %d \n",index ,numbers)
	}

	//to print the index only
	for index:=range numbers{
		fmt.Printf("The index is %d \n",index)
	}

	//to print number only we use _ 
	for _, numbers := range numbers{
		fmt.Printf("The number is %d \n",numbers)
	}
}

