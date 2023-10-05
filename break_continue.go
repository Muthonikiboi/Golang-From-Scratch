//Go Break and Continue statements

package main
import "fmt"

func main(){
	//break
	for i:=0 ; i<10 ; i++ {
		if i==5{
			break
		}
		fmt.Println(i)
	}
	//Go continue =the value is skipped and continues to another loop

	for j:=0 ; j<10 ; j++ {
		if j==5{
			fmt.Println("The value of 5 will not be included")
			continue
		}
		fmt.Println(j)
	}
}