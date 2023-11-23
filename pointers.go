package main
import "fmt"

func main() {
  
  var num int = 5
  fmt.Println("Variable Value:", num)

  fmt.Println("Memory Address:", &num)

  //using pointer
  var name = "Joy Kiboi"
  var ptr *string

  ptr = &name

  fmt.Println("Value of pointer is", ptr)
  fmt.Println("Address of the variable", &name)

}