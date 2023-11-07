package main
import "fmt"

func greet() func() {

  return func() {
    fmt.Println("My name is Joy")
  }

}

func main() {

  g1 := greet()
  g1()
}