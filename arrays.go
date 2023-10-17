package main
 
import "fmt"
 
func main() {
  var arry [5] string
  fmt.Println(arry)

  var num[5] int
  num[1]=500
  num[4]=670
  fmt.Println(num)
  fmt.Println(num[4])

  // second method to define an array
  arr :=[3] string {"GO","C#","Javascript"}
  
  //get length of array in go
  length := len(arr)
  fmt.Println(length)

  fmt.Println(arr)

  //code to loop through an array
  for i:=0; i<len(arr); i++ {
    fmt.Println(arr[i])
  }
}