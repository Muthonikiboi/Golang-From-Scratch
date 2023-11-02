package main
import "fmt"

func countDown(number int) {

  if number > 0 {
    fmt.Println(number)

    countDown(number - 1)
  } else {
    fmt.Println("Countdown Stops")
  }
 }
  //example of recursion using factorial ( )

  func factorial(n int) int {
    if n<=1{
      return 1
    }
    return n * factorial(n-1)
  }
func main() {
  countDown(3)
    fmt.Println(factorial(5))
  }