//this is changing the value of one datatypee to another
package main
import "fmt"

func main(){
	//Exlicitly
	//change float to integer
	var float1 float32=5.22
	var integer1 int= int(float1)
	fmt.Printf("The value of float is :%g \n", float1)
	fmt.Printf("The value of integer is : %d", integer1)

	//change integer to float
	var integer2 int=56
	var float2 float32 =float32(integer2)
	fmt.Printf("The value of integer is :%d \n", integer2)
	fmt.Printf("The value of float is :%g :",float2)
	
}